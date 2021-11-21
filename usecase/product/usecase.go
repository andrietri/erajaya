package product

import (
	"context"
	"erajaya/entities/model"
	repository "erajaya/repository/product"
)

type (
	ProductUsecase interface {
		Fetch(context.Context, string, string) ([]model.ProductResponse, error)
		Store(context.Context, model.ProductRequest) (*model.ProductResponse, error)
	}

	productUsecase struct {
		productRepository repository.ProductRepository
	}
)

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}
