package middleware

import (
	"net"
	"net/http"
	"strings"
	"time"

	_ "embed"

	"rate-limiter/config"
	"rate-limiter/internal/redis"
)

//go:embed token_bucket.lua
var luaScript string


func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Extract client IP
		ip := r.RemoteAddr
		if strings.Contains(ip, ":") {
			ip, _, _ = net.SplitHostPort(ip)
		}

		// Build Redis key
		key := config.RateLimitKeyPrefix + ip

		// Current timestamp
		now := time.Now().Unix()

		// Execute Lua script atomically
		allowed, err := redis.Client.Eval(
			redis.Ctx,
			luaScript,
			[]string{key},
			config.BucketCapacity,
			config.RefillRatePerSecond,
			now,
		).Int()

		// Block if rate limited
		if err != nil || allowed == 0 {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("429 Too Many Requests\n"))
			return
		}

		// Allow request
		next.ServeHTTP(w, r)
	})
}




