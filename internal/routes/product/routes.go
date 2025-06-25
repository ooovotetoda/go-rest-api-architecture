package user

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	productsHandlers "go-rest-api-architecture/internal/handlers/products"
	productsRepository "go-rest-api-architecture/internal/repositories/products"
	productsServices "go-rest-api-architecture/internal/services/products"
	"log/slog"
)

func RegisterRoutes(r *chi.Mux, ctx context.Context, log *slog.Logger, db *pgxpool.Pool) {
	productsRepository := productsRepository.NewProductsRepository(db)
	productsService := productsServices.NewProductsService(productsRepository)
	productsHandlers := productsHandlers.NewProductsHandlers(log, productsService)

	r.Route("/products", func(r chi.Router) {
		r.Get("/{id}", productsHandlers.GetProduct(ctx))
	})
}
