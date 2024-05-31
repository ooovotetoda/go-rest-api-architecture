package users

import (
	resp "cabinet-wooffie-api/internal/lib/api/response"
	"cabinet-wooffie-api/internal/lib/logger/sl"
	"cabinet-wooffie-api/internal/services/users"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"strconv"
)

type UsersHandlers struct {
	s   *users.UsersService
	log *slog.Logger
}

func NewUsersHandlers(log *slog.Logger, service *users.UsersService) *UsersHandlers {
	return &UsersHandlers{
		s:   service,
		log: log,
	}
}

func (h *UsersHandlers) GetUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.users.GetUser"

		log := h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id := chi.URLParam(r, "id")
		userID, err := strconv.Atoi(id)
		if err != nil {
			log.Error("Failed to convert users ID", sl.Err(err))
			render.JSON(w, r, resp.Error("internal error"))
			return
		}

		user, err := h.s.GetUser(ctx, int32(userID))
		if err != nil {
			log.Error("Failed to get user by ID", sl.Err(err))
			render.JSON(w, r, resp.Error("internal error"))
			return
		}

		render.JSON(w, r, user)
		log.Info("User fetched successfully")
	}
}
