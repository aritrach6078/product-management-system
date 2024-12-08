package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client // Exported Redis client variable

func InitializeRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis address
		DB:   0,                // Default DB
	})

	ctx := context.Background()
	if _, err := Rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Connected to Redis successfully!")
}
