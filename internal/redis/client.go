package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"rate-limiter/config"
)

var (
	Ctx    = context.Background()
	Client *redis.Client
)

// Init initializes the Redis client
func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr: config.RedisAddress,
		  Username: config.RedisUsername,
        Password: config.RedisPassword,
	})

	// Check connection
	if err := Client.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
}
