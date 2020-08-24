package main

import (
	"fmt"
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"

	"github.com/lukewhrit/lynnx/config"
	"github.com/lukewhrit/lynnx/database"
	"github.com/lukewhrit/lynnx/routes"
)

func main() {
	// Load configuration
	if err := config.Load(); err != nil {
		log.Fatalf("Couldn't load configuration file: %v", err)
	}

	// Initialize app
	app := fiber.New()

	// Register middleware and endpoints
	registerMiddleware(app)
	routes.Register(app)
	app.Static("/", "./static")

	database.Load()

	// Start the server
	address := fmt.Sprintf("%s:%d", config.Configuration.Server.Host, config.Configuration.Server.Port)

	log.Fatal(app.Listen(address))
}

func registerMiddleware(app *fiber.App) {
	app.Use(middleware.Compress(middleware.CompressConfig{
		Level: config.Configuration.CompressionLevel,
	}))

	app.Use(cors.New())
	app.Use(middleware.Logger())

	app.Use(func(c *fiber.Ctx) {
		c.Set("X-Download-Options", "noopen")
		c.Set("X-DNS-Prefetch-Control", "off")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("Referrer-Policy", "no-referrer-when-downgrade")
		c.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")

		if config.Configuration.Server.EnableCSP == true {
			c.Set("Content-Security-Policy", "default-src 'self' https:; frame-ancestors 'none'; base-uri 'none'; form-action 'self' https:; img-src data:;")
		}

		c.Next()
	})
}
