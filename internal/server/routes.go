package server

import (
	productsHandlers "cabinet-wooffie-api/internal/handlers/products"
	usersHandlers "cabinet-wooffie-api/internal/handlers/users"
	mwCors "cabinet-wooffie-api/internal/middleware/cors"
	productsRepository "cabinet-wooffie-api/internal/repositories/products"
	usersRepository "cabinet-wooffie-api/internal/repositories/users"
	productsServices "cabinet-wooffie-api/internal/services/products"
	usersServices "cabinet-wooffie-api/internal/services/users"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
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

	usersRepository := usersRepository.NewUsersRepository(db)
	productsRepository := productsRepository.NewProductsRepository(db)

	usersService := usersServices.NewUsersService(usersRepository)
	productsService := productsServices.NewProductsService(productsRepository)

	usersHandlers := usersHandlers.NewUsersHandlers(log, usersService)
	productsHandlers := productsHandlers.NewProductsHandlers(log, productsService)

	router.Get("/users/{id}", usersHandlers.GetUser(ctx))
	router.Get("/products/{id}", productsHandlers.GetProduct(ctx))

	return router
}
