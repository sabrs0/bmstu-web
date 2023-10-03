package main

import (
	"os"

	"log/slog"

	cfg "github.com/sabrs0/bmstu-web/src/internal/config"
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
	_ = storage
	//router
	//server
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
