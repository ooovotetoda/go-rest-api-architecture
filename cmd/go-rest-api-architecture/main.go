package main

import (
	"context"
	"go-rest-api-architecture/internal/configs"
	"go-rest-api-architecture/internal/database/postgres"
	setupLog "go-rest-api-architecture/internal/lib/logger/setup"
	"go-rest-api-architecture/internal/lib/logger/sl"
	"go-rest-api-architecture/internal/server"
	"log/slog"
	"os"
)

func main() {
	cfg := configs.MustLoad()

	log := setupLog.SetupLogger(cfg.Env)

	log.Info("App started", slog.String("env", cfg.Env))
	log.Debug("Debugging started")

	db, err := postgres.New(cfg)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	defer db.Close()

	ctx := context.Background()
	server := server.NewServer(ctx, cfg, log, db)

	err = server.ListenAndServe()
	if err != nil {
		log.Error("failed to start server", sl.Err(err))
	}
}
