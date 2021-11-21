package product

import (
	"context"
	"erajaya/entities/schema"
)

func (m *productRepository) Store(ctx context.Context, data schema.Product) (*schema.Product, error) {
	result := m.Conn.WithContext(ctx).Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}
