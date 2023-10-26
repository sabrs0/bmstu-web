package main

import (
	"net/http"
	"os"

	"log/slog"

	cfg "github.com/sabrs0/bmstu-web/src/internal/config"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/routers/gorilla"
	"github.com/sabrs0/bmstu-web/src/internal/lib/logger/sl"
	"github.com/sabrs0/bmstu-web/src/internal/storage/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := cfg.MustLoad() //config

	log := setupLogger(cfg.Env) //logger
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("Debug msgs are enabled")

	storage, err := postgres.New() //storage
	if err != nil {
		log.Error("Failed to init database", sl.Err(err))
	}
	router := gorilla.SetRouter(storage.DB, log)
	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.HTTPServer.Address,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
	}
	log.Info("Starting server at ",
		slog.String("addr", srv.Addr))
	if err := srv.ListenAndServe(); err != nil { //server
		log.Error("failed to start server")
	}
	log.Error("server stopped")

}
func setupLogger(env string) *slog.Logger {
	var logger *slog.Logger
	switch env {
	case envLocal:
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return logger

}
