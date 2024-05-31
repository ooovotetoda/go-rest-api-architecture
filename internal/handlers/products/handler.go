package products

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	resp "go-rest-api-architecture/internal/lib/api/response"
	"go-rest-api-architecture/internal/lib/logger/sl"
	"go-rest-api-architecture/internal/services/products"
	"log/slog"
	"net/http"
	"strconv"
)

type ProductsHandlers struct {
	s   *products.ProductsService
	log *slog.Logger
}

func NewProductsHandlers(log *slog.Logger, service *products.ProductsService) *ProductsHandlers {
	return &ProductsHandlers{
		s:   service,
		log: log,
	}
}

func (h *ProductsHandlers) GetProduct(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.products.GetProduct"

		log := h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id := chi.URLParam(r, "id")
		productID, err := strconv.Atoi(id)
		if err != nil {
			log.Error("Failed to convert products ID", sl.Err(err))
			render.JSON(w, r, resp.Error("internal error"))
			return
		}

		product, err := h.s.GetProduct(int32(productID))
		if err != nil {
			log.Error("Failed to get product by ID", sl.Err(err))
			render.JSON(w, r, resp.Error("internal error"))
			return
		}

		render.JSON(w, r, product)
		log.Info("Product fetched successfully")
	}
}
