package cache

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/renatafborges/rate-limiter-go/configs"
)

var (
	redisClient *redis.Client
)

type limiter struct {
}

func LoadEnvCache() {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
}

func NewCheckRateLimit() CacheInterface {
	return &limiter{}

}

func (l *limiter) CheckRateLimit(key string, limit int) (bool, error) {

	ctx := context.Background()

	val, err := redisClient.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		slog.Error("unable to retrieve values from redis client", "key:", key, "ctx:", ctx, "error:", err)
		return false, err
	}

	if val >= limit {
		return false, nil
	}

	_, err = redisClient.Incr(ctx, key).Result()
	if err != nil {
		slog.Error("unable to increment values to redis client", "key:", key, "ctx:", ctx, "error:", err)
		return false, err
	}

	if val == 0 {
		err = redisClient.Expire(ctx, key, configs.BlockDuration).Err()
		if err != nil {
			slog.Error("unable to set an expiration time", "key:", key, "ctx:", ctx, "error:", err)
			return false, err
		}
	}

	return true, nil
}
