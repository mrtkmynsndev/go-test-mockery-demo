package services

import (
	"errors"

	"github.com/mrtkmynsndev/go-test-mockery-demo/internal/models"
	"github.com/mrtkmynsndev/go-test-mockery-demo/internal/repositories"
)

type ProductService struct {
	repo repositories.ProductRepositoryInterface
}

func NewProductService(repo repositories.ProductRepositoryInterface) ProductService {
	return ProductService{
		repo: repo,
	}
}

func (s ProductService) Insert(productID string, product models.InsertProductDTO) error {
	if len(productID) == 0 {
		return errors.New("productID can not be null")
	}

	err := s.repo.Add(models.Product{
		ID:    productID,
		Name:  product.Name,
		Price: product.Price,
		Stock: product.Stock,
	})
	if err != nil {
		return err
	}

	return nil
}
