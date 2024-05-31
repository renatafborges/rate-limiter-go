package cache

type CacheInterface interface {
	CheckRateLimit(key string, limit int) (bool, error)
}
