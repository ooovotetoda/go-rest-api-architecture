package postgres

import (
	"cabinet-wooffie-api/internal/configs"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func New(cfg *configs.Config) (*pgxpool.Pool, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s users=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open SQL connection: %w", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to check SQL connection: %w", err)
	}

	//if err := goose.Up(db, "migrations"); err != nil {
	//	return nil, fmt.Errorf("failed to run migrations: %w", err)
	//}

	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, psqlInfo)
	if err != nil {
		return nil, err
	}

	return dbpool, nil
}
