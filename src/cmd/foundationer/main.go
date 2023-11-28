//   Version: 0.1
//
//	Schemes: http
//	Host: localhost:8081
//	BasePath: /api/v1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//
//	SecurityDefinitions:
//	bearerAuth:
//	    type: apiKey
//	    name: Authorization
//	    in: header
//
// swagger:meta
package main

import (
	"net/http"
	"os"
	"time"

	"log/slog"

	"github.com/rs/cors"
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
	log.Info("starting foundationer", slog.String("env", cfg.Env))
	log.Debug("Debug msgs are enabled")
	time.Sleep(time.Second * 15)
	storage, err := postgres.New() //storage
	if err != nil {
		log.Error("Failed to init database", sl.Err(err))
	}
	//log.Info("Successfully connected to database")
	router := gorilla.SetRouter(storage.DB, log)

	corsHandler := cors.Default().Handler(router)
	srv := &http.Server{
		Handler:      corsHandler,
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
