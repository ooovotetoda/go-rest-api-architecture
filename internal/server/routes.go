package server

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	productsHandlers "go-rest-api-architecture/internal/handlers/products"
	usersHandlers "go-rest-api-architecture/internal/handlers/users"
	mwCors "go-rest-api-architecture/internal/middleware/cors"
	productsRepository "go-rest-api-architecture/internal/repositories/products"
	usersRepository "go-rest-api-architecture/internal/repositories/users"
	productsServices "go-rest-api-architecture/internal/services/products"
	usersServices "go-rest-api-architecture/internal/services/users"
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
