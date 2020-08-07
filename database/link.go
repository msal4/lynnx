package database

import (
	"github.com/lukewhrit/lynnx/config"
	"github.com/lukewhrit/lynnx/utils"
)

// NewLink inserts into the database
func NewLink(value string) (string, error) {
	key := utils.RandomString(config.Configuration.LinkLength)
	err := RedisClient.Set(Ctx, key, value, 0).Err()

	return key, err
}

// GetLink gets values from the database via their key
func GetLink(key string) (string, error) {
	return RedisClient.Get(Ctx, key).Result()
}
