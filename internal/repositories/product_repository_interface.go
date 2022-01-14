package repositories

import "github.com/mrtkmynsndev/go-test-mockery-demo/internal/models"

type ProductRepositoryInterface interface {
	Add(product models.Product) error
}
