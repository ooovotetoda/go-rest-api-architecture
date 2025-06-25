package user

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	usersHandlers "go-rest-api-architecture/internal/handlers/users"
	usersRepository "go-rest-api-architecture/internal/repositories/users"
	usersServices "go-rest-api-architecture/internal/services/users"
	"log/slog"
)

func RegisterRoutes(r *chi.Mux, ctx context.Context, log *slog.Logger, db *pgxpool.Pool) {
	usersRepository := usersRepository.NewUsersRepository(db)
	usersService := usersServices.NewUsersService(usersRepository)
	usersHandlers := usersHandlers.NewUsersHandlers(log, usersService)

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", usersHandlers.GetUser(ctx))
	})
}
