package products

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-rest-api-architecture/internal/database/sqlc"
	"go-rest-api-architecture/internal/models"
	"log"
)

type ProductsRepository struct {
	db *pgxpool.Pool
	q  *sqlc.Queries
}

func NewProductsRepository(db *pgxpool.Pool) *ProductsRepository {
	return &ProductsRepository{
		db: db,
		q:  sqlc.New(db),
	}
}

func (r *ProductsRepository) GetProduct(id int32) (*models.Product, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, args, err := psql.Select("id", "name", "description", "price").
		From("products").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		log.Fatalf("Failed to build query: %v\n", err)
	}

	var product models.Product
	err = r.db.QueryRow(context.Background(), query, args...).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		log.Fatalf("Failed to execute query: %v\n", err)
	}

	return &product, nil
}
