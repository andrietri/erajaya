package product

import (
	"context"
	"erajaya/cache"
	"erajaya/entities/model"
	"erajaya/entities/schema"
)

func (p *productUsecase) Store(ctx context.Context, data model.ProductRequest) (*model.ProductResponse, error) {
	request := schema.Product{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Quantity:    data.Quantity,
	}

	res, err := p.productRepository.Store(ctx, request)
	if err != nil {
		return nil, err
	}

	// start delete cache
	redisClient := cache.NewClient()
	iter := redisClient.Scan(0, "products_*", 0).Iterator()
	for iter.Next() {
		err := redisClient.Del(iter.Val()).Err()
		if err != nil {
			panic(err)
		}
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
	// stop delete cache

	product := model.ProductResponse{
		Id:          int(res.Id),
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		Quantity:    res.Quantity,
	}

	return &product, nil
}
