package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/lukewhrit/lynnx/config"
)

var (
	// Ctx is a thing, I guess. I honestly don't know
	Ctx = context.Background()

	// RedisClient holds the active connection to the database
	RedisClient *redis.Client
)

// Load creates a connection to the database
func Load() {
	opts, err := redis.ParseURL(config.Configuration.RedisURI)

	if err != nil {
		log.Fatalf("Couldn't parse Redis URI: %s", err)
	}

	RedisClient = redis.NewClient(opts)
}
