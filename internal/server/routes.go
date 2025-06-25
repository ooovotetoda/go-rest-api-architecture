package server

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	mwCors "go-rest-api-architecture/internal/middleware/cors"
	product "go-rest-api-architecture/internal/server/routes/product"
	user "go-rest-api-architecture/internal/server/routes/user"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes(ctx context.Context, log *slog.Logger, db *pgxpool.Pool) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	//router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Use(mwCors.New())

	user.RegisterRoutes(router, ctx, log, db)
	product.RegisterRoutes(router, ctx, log, db)

	return router
}
