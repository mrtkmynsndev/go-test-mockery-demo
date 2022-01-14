package services_test

import (
	"errors"
	"testing"

	"github.com/mrtkmynsndev/go-test-mockery-demo/internal/mocks/repomocks"
	"github.com/mrtkmynsndev/go-test-mockery-demo/internal/models"
	"github.com/mrtkmynsndev/go-test-mockery-demo/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	onErr bool
}

func (r *MockProductRepository) Add(models.Product) error {
	if r.onErr {
		return errors.New("error while inserting")
	}
	return nil
}

func TestProductService_Insert(t *testing.T) {
	repo := &repomocks.ProductRepositoryInterface{}
	repo.On("Add", mock.AnythingOfType("models.Product")).
		Return(nil).
		Once()

	service := services.NewProductService(repo)

	err := service.Insert("2f1afe98-63c4-4f59-bcaf-1df835602bdb", models.InsertProductDTO{
		Name:  "Macbook",
		Price: 20500,
		Stock: 10,
	})

	assert.Nil(t, err)
}
