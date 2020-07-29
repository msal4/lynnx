package database

import (
	"github.com/lukewhrit/lynnx/config"
	"github.com/lukewhrit/lynnx/utils"
)

// Create inserts into the database
func Create(value string) (string, error) {
	key := utils.RandomString(config.GetConfig().LinkLength)
	err := RedisClient.Set(Ctx, key, value, 0).Err()

	return key, err
}

// Read gets values from the database via their key
func Read(key string) (string, error) {
	return RedisClient.Get(Ctx, key).Result()
}
