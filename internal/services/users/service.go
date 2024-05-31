package users

import (
	"cabinet-wooffie-api/internal/database/sqlc"
	"cabinet-wooffie-api/internal/repositories/users"
	"context"
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
