package product

import (
	"context"
	"encoding/json"
	"erajaya/cache"
	"erajaya/entities/model"
	"fmt"
	"time"
)

func (p *productUsecase) Fetch(ctx context.Context, sortBy string, sortType string) ([]model.ProductResponse, error) {
	if sortBy == "" {
		sortBy = "created_at"
	}

	if sortType == "" {
		sortType = "desc"
	}

	results := []model.ProductResponse{}

	// start search value in redis
	redisClient := cache.NewClient()
	nameInRedis, err := redisClient.Get("products_all").Result()
	// stop search value in redis

	if err == nil {
		err = json.Unmarshal([]byte(nameInRedis), &results)
		if err != nil {
			panic(err)
		}

		return results, nil
	}

	res, err := p.productRepository.Fetch(ctx, sortBy, sortType)
	if err != nil {
		return nil, err
	}

	for i := range res {
		rest := model.ProductResponse{
			Id:          int(res[i].Id),
			Name:        res[i].Name,
			Description: res[i].Description,
			Price:       res[i].Price,
			Quantity:    res[i].Quantity,
			CreatedAt:   fmt.Sprintf("%v", res[i].CreatedAt),
		}

		results = append(results, rest)
	}

	json, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}
	_ = redisClient.Set("products_all", json, 120*time.Second)

	return results, nil
}
