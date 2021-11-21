package product

import (
	"context"
	"erajaya/entities/schema"
	"errors"
	"fmt"
)

func (m *productRepository) Fetch(ctx context.Context, sortBy string, sortType string) ([]schema.Product, error) {
	order := fmt.Sprintf("%s %s", sortBy, sortType)

	products := []schema.Product{}
	result := m.Conn.Order(order).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(products) == 0 {
		return nil, errors.New("data not found")
	}
	return products, nil
}
