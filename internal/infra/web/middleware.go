package web

import (
	"log/slog"
	"net/http"

	"github.com/renatafborges/rate-limiter-go/configs"
	"github.com/renatafborges/rate-limiter-go/internal/infra/cache"
)

func RateLimiter(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ip := r.RemoteAddr
		token := r.Header.Get("API_KEY")

		key := ip
		limit := configs.RateLimitIP

		if token != "" {
			key = token
			limit = configs.RateLimitToken
		}

		limiter := cache.NewCheckRateLimit()

		allowed, err := limiter.CheckRateLimit(key, limit)
		if err != nil {
			slog.Error("unable to check rate limit", "key:", key, "limit:", limit, "error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if !allowed {
			http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
