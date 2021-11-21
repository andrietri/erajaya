package cache

import (
	"os"

	"github.com/go-redis/redis/v7"
)

func NewClient() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: "", // no password
		DB:       0,  // default DB
	})

	return redisClient
}

func Ping(client *redis.Client) (string, error) {
	result, err := client.Ping().Result()

	if err != nil {
		return "", err
	} else {
		return result, nil
	}
}
