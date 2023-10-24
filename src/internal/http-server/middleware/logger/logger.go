package logger

import (
	"log/slog"
	"net/http"
	"time"
)

// принимает созданный логгер
// Возвращает мидлварную функцию.
func New(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		logger := logger.With("component", "middleware/logger")
		logger.Info("logger middleware enabled")
		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := logger.With(
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
			)
			//ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			startReqTime := time.Now()
			//логгирование выполнится после обработки запроса
			defer func() {
				entry.Info("Request comleted",
					slog.String("duration", time.Since(startReqTime).String()),
					//slog.Int("status", r.Response.StatusCode), //ww
				)
			}()

			next.ServeHTTP(w, r) //ww
		}
		return http.HandlerFunc(fn)
	}
}
