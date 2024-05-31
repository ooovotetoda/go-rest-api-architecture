package server

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"go-rest-api-architecture/internal/configs"
	"log/slog"
	"net/http"
)

type Server struct {
	address string
	log     *slog.Logger
	db      *pgxpool.Pool
}

func NewServer(ctx context.Context, cfg *configs.Config, log *slog.Logger, db *pgxpool.Pool) *http.Server {
	NewServer := &Server{
		address: cfg.Server.Address,
		log:     log,
		db:      db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      NewServer.RegisterRoutes(ctx, log, db),
		IdleTimeout:  cfg.Server.IdleTimeout,
		ReadTimeout:  cfg.Server.Timeout,
		WriteTimeout: cfg.Server.Timeout,
	}

	return server
}
