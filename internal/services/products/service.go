package products

import (
	"cabinet-wooffie-api/internal/models"
	"cabinet-wooffie-api/internal/repositories/products"
)

type ProductsService struct {
	r *products.ProductsRepository
}

func NewProductsService(r *products.ProductsRepository) *ProductsService {
	return &ProductsService{
		r: r,
	}
}

func (s *ProductsService) GetProduct(id int32) (*models.Product, error) {
	product, err := s.r.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
