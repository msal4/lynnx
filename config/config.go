package config

import (
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

var k = koanf.New(".")

// Server is the configuration object for anything server-related
type Server struct {
	Port             int    `koanf:"port"`
	Host             string `koanf:"host"`
	CompressionLevel int    `koanf:"compression_level"`
	EnableCSP        bool   `koanf:"enable_csp"`
}

// Config is the configuration object
type Config struct {
	Server `koanf:"server"`

	LinkLength      int    `koanf:"link_length"`
	AllowCustomURLs bool   `koanf:"allow_custom_urls"`
	RedisURI        string `koanf:"redis_uri"`
}

var configuration Config

// Load configuration from file
func Load() error {
	// Set some default values
	k.Load(confmap.Provider(map[string]interface{}{
		"server.port":              3000,
		"server.host":              "0.0.0.0",
		"server.compression_level": -1,
		"server.enable_csp":        true,
		"link_length":              8,
		"allow_custom_urls":        false,
		"redis_uri":                "redis://localhost:6379/0",
	}, "."), nil)

	f := file.Provider("./config.json")

	// Load configuration from JSON on top of said default values
	if err := k.Load(f, json.Parser()); err != nil {
		log.Fatalf("error loading config from file: %v", err)
	}

	err := k.Load(env.Provider("LYNNX_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "LYNNX_")), "_", ".", -1)
	}), nil)

	if err != nil {
		log.Fatalf("error loading config from env: %v", err)
	}

	k.Unmarshal("", &configuration)

	return nil
}

// GetConfig returns the entire configuration object
func GetConfig() Config {
	return configuration
}
