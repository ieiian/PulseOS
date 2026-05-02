package middleware

import (
	"net/http"
	"time"

	"github.com/tse/PulseOS/backend/internal/pkg/logger"
)

func WithAccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Infof("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}
