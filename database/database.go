package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/lukewhrit/lynnx/config"
)

var (
	ctx = context.Background()

	// RedisClient holds the active connection to the database
	RedisClient *redis.Client
)

// Load creates a connection to the database
func Load() {
	opts, err := redis.ParseURL(config.GetConfig().RedisURI)

	if err != nil {
		log.Fatalf("Couldn't parse Redis URI: %s", err)
	}

	RedisClient = redis.NewClient(opts)

	pong, err := RedisClient.Ping(ctx).Result()
	fmt.Println(pong, err)
}
