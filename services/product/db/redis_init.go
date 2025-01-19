package db

import (
	"context"
	"go-microservices/products/config"

	"log"

	"github.com/redis/go-redis/v9"
)

var REDIS_CLIENT *redis.Client

func InitRedis(SERVICE_CONFIG config.Config) {
	REDIS_CLIENT = redis.NewClient(&redis.Options{
		Addr: SERVICE_CONFIG.Redis_URL, // Adjust host and port as needed
		DB:   0,                        // Default DB
	})

	if err := REDIS_CLIENT.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}
