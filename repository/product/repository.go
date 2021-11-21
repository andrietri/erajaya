package product

import (
	"context"
	"erajaya/entities/schema"

	"gorm.io/gorm"
)

type (
	ProductRepository interface {
		Fetch(context.Context, string, string) ([]schema.Product, error)
		Store(context.Context, schema.Product) (*schema.Product, error)
	}

	productRepository struct {
		Conn *gorm.DB
	}
)

func NewProductRepository(dbCon *gorm.DB) ProductRepository {
	return &productRepository{
		Conn: dbCon,
	}
}
