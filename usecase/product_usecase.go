package usecase

import (
	"github.com/Senhordim/go-api/model"
	"github.com/Senhordim/go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	products, err := pu.repository.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	id, error := pu.repository.CreateProduct(product)

	if error != nil {
		return model.Product{}, error
	}

	product.ID = id

	return product, nil
}
