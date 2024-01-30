package metrics

import (
	"net/http"
	"time"

	"github.com/sabrs0/bmstu-web/src/test/benchmark"
)

func New(metrics *benchmark.Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			metrics.Timings.WithLabelValues(r.URL.Path).Observe(float64(time.Since(start).Seconds()))
			metrics.Counter.WithLabelValues(r.URL.Path).Inc()
		})
	}
}
