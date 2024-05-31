package users

import (
	"context"
	"go-rest-api-architecture/internal/database/sqlc"
	"go-rest-api-architecture/internal/repositories/users"
)

type UsersService struct {
	r *users.UsersRepository
}

func NewUsersService(r *users.UsersRepository) *UsersService {
	return &UsersService{
		r: r,
	}
}

func (s *UsersService) GetUser(ctx context.Context, id int32) (*sqlc.User, error) {
	user, err := s.r.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
