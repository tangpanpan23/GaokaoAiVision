package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type RateLimitMiddleware struct {
	requests map[string][]time.Time
	mu       sync.RWMutex
	limits   map[string]int // 不同时间窗口的限制
}

func NewRateLimitMiddleware() *RateLimitMiddleware {
	return &RateLimitMiddleware{
		requests: make(map[string][]time.Time),
		limits: map[string]int{
			"second": 100,  // 每秒最多100个请求
			"minute": 1000, // 每分钟最多1000个请求
			"hour":   10000, // 每小时最多10000个请求
		},
	}
}

func (m *RateLimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientIP := getClientIP(r)

		m.mu.Lock()
		now := time.Now()

		// 清理过期请求记录
		m.cleanup(clientIP, now)

		// 检查各种时间窗口的限流
		if m.isRateLimited(clientIP, now) {
			m.mu.Unlock()
			httpx.ErrorCtx(r.Context(), w, &struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: 429,
				Msg:  "请求过于频繁，请稍后再试",
			})
			return
		}

		// 记录请求
		m.requests[clientIP] = append(m.requests[clientIP], now)
		m.mu.Unlock()

		next(w, r)
	}
}

func (m *RateLimitMiddleware) isRateLimited(clientIP string, now time.Time) bool {
	requests := m.requests[clientIP]

	// 检查每秒限流
	secondCount := 0
	for _, reqTime := range requests {
		if now.Sub(reqTime) < time.Second {
			secondCount++
		}
	}
	if secondCount >= m.limits["second"] {
		return true
	}

	// 检查每分钟限流
	minuteCount := 0
	for _, reqTime := range requests {
		if now.Sub(reqTime) < time.Minute {
			minuteCount++
		}
	}
	if minuteCount >= m.limits["minute"] {
		return true
	}

	// 检查每小时限流
	hourCount := 0
	for _, reqTime := range requests {
		if now.Sub(reqTime) < time.Hour {
			hourCount++
		}
	}
	if hourCount >= m.limits["hour"] {
		return true
	}

	return false
}

func (m *RateLimitMiddleware) cleanup(clientIP string, now time.Time) {
	if requests, exists := m.requests[clientIP]; exists {
		var validRequests []time.Time
		for _, reqTime := range requests {
			if now.Sub(reqTime) < time.Hour {
				validRequests = append(validRequests, reqTime)
			}
		}
		if len(validRequests) == 0 {
			delete(m.requests, clientIP)
		} else {
			m.requests[clientIP] = validRequests
		}
	}
}

func getClientIP(r *http.Request) string {
	// 优先获取X-Forwarded-For头
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		// X-Forwarded-For可能包含多个IP，取第一个
		return strings.Split(xForwardedFor, ",")[0]
	}

	// 其次获取X-Real-IP头
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		return xRealIP
	}

	// 最后获取RemoteAddr
	return r.RemoteAddr
}

