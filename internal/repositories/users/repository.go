package users

import (
	"cabinet-wooffie-api/internal/database/sqlc"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	db *pgxpool.Pool
	q  *sqlc.Queries
}

func NewUsersRepository(db *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{
		db: db,
		q:  sqlc.New(db),
	}
}

func (r *UsersRepository) GetUser(ctx context.Context, id int32) (*sqlc.User, error) {
	user, err := r.q.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
