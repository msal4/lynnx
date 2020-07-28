package config

import (
	"github.com/spf13/viper"
)

// Server is the configuration object for anything server-related
type Server struct {
	Port             int
	Host             string
	CompressionLevel int
	EnableCSP        bool
}

// Config is the configuration object
type Config struct {
	Server

	LinkLength      int
	AllowCustomURLs bool
	RedisURI        string
}

var configuration *Config

// Load configuration from file
func Load() error {
	c := new(Config)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.SetDefault("Server.Port", 3000)
	viper.SetDefault("Server.Host", "0.0.0.0")
	viper.SetDefault("Server.CompressionLevel", 1)
	viper.SetDefault("Server.EnableCSP", true)
	viper.SetDefault("LinkLength", 8)
	viper.SetDefault("AllowCustomURLs", false)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	configuration = c

	return nil
}

// GetConfig returns the entire configuration object
func GetConfig() Config {
	return *configuration
}
