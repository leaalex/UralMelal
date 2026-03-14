package middleware

import (
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	mu       sync.Mutex
	visitors map[string]*visitor
	limit    int
	window   time.Duration
}

type visitor struct {
	count    int
	windowAt time.Time
}

func RateLimit(limit int, window time.Duration) func(http.Handler) http.Handler {
	rl := &rateLimiter{
		visitors: make(map[string]*visitor),
		limit:    limit,
		window:   window,
	}
	go rl.cleanup()
	return rl.handler
}

func (rl *rateLimiter) handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.RemoteAddr
		if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
			key = fwd
		}
		rl.mu.Lock()
		v, ok := rl.visitors[key]
		if !ok || time.Since(v.windowAt) > rl.window {
			v = &visitor{count: 0, windowAt: time.Now()}
			rl.visitors[key] = v
		}
		v.count++
		allowed := v.count <= rl.limit
		rl.mu.Unlock()
		if !allowed {
			http.Error(w, `{"error":"too many requests"}`, http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (rl *rateLimiter) cleanup() {
	for range time.Tick(time.Minute) {
		rl.mu.Lock()
		now := time.Now()
		for k, v := range rl.visitors {
			if now.Sub(v.windowAt) > rl.window {
				delete(rl.visitors, k)
			}
		}
		rl.mu.Unlock()
	}
}
