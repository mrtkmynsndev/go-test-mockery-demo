package repositories

import (
	"sync"

	"github.com/mrtkmynsndev/go-test-mockery-demo/internal/models"
)

type ProductRepository struct {
	data  map[string]models.Product
	mutex sync.Mutex
}

func NewProductRepository() ProductRepository {
	return ProductRepository{
		data:  make(map[string]models.Product),
		mutex: sync.Mutex{},
	}
}

func (r *ProductRepository) Add(product models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.data[product.ID] = product
	return nil
}
