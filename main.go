package main

import (
	"log"

	"github.com/mrtkmynsndev/go-test-mockery-demo/internal/models"
	"github.com/mrtkmynsndev/go-test-mockery-demo/internal/repositories"
	"github.com/mrtkmynsndev/go-test-mockery-demo/internal/services"
)

func main() {
	//creating dependency
	repo := repositories.NewProductRepository()
	service := services.NewProductService(&repo)

	// inserting
	err := service.Insert("d681753c-dca3-433d-9d8a-0d455d4a3065", models.InsertProductDTO{
		Name:  "Macbook",
		Price: 27500,
		Stock: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	// inserting
	err = service.Insert("6220d651-2c62-43b9-bf31-2a6f79059f69", models.InsertProductDTO{
		Name:  "HP Laptop",
		Price: 13500,
		Stock: 5,
	})
	if err != nil {
		log.Fatal(err)
	}

	// inserting
	err = service.Insert("77076fae-a4a9-4e73-ad21-7fb89f4ab39d", models.InsertProductDTO{
		Name:  "Iphone 13 pro",
		Price: 24500,
		Stock: 25,
	})
	if err != nil {
		log.Fatal(err)
	}
}
