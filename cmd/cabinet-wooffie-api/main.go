package main

import (
	"cabinet-wooffie-api/internal/configs"
	"cabinet-wooffie-api/internal/database/postgres"
	setupLog "cabinet-wooffie-api/internal/lib/logger/setup"
	"cabinet-wooffie-api/internal/lib/logger/sl"
	"cabinet-wooffie-api/internal/server"
	"context"
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
