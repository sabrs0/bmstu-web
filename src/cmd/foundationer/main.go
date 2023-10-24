package main

import (
	"net/http"
	"os"

	"log/slog"

	"github.com/gorilla/mux"
	cfg "github.com/sabrs0/bmstu-web/src/internal/config"
	ctrl "github.com/sabrs0/bmstu-web/src/internal/controllers"
	repos "github.com/sabrs0/bmstu-web/src/internal/dataAccess/repositories/postgres"
	handlers "github.com/sabrs0/bmstu-web/src/internal/http-server/handlers/foundations"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/middleware/logger"
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
	router := mux.NewRouter() //router
	//добавить мидлвар для проверки авторизации
	router.Use(logger.New(log))
	foundationRepo := repos.NewFoundationRepository(storage.DB)
	foundrisingRepo := repos.NewFoundrisingRepository(storage.DB)
	transationRepo := repos.NewTransactionRepository(storage.DB)
	foundationController := ctrl.NewFoundationController(foundationRepo, transationRepo, foundrisingRepo)
	router.HandleFunc("/foundations", handlers.GetAll(log, foundationController)).Methods(http.MethodGet)
	router.HandleFunc("/foundations", handlers.Add(log, foundationController)).Methods(http.MethodPost)
	srv := &http.Server{
		Handler:      router,
		Addr:         cfg.HTTPServer.Address,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
	}
	log.Info("Starting server at ",
		slog.String("addr", srv.Addr))
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
	log.Error("server stopped")
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
