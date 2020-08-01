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

	"github.com/lukewhrit/middlewares"
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
	address := fmt.Sprintf("%s:%d", config.GetConfig().Server.Host, config.GetConfig().Server.Port)

	log.Fatal(app.Listen(address))
}

func registerMiddleware(app *fiber.App) {
	app.Use(middleware.Compress(middleware.CompressConfig{
		Level: config.GetConfig().CompressionLevel,
	}))

	app.Use(cors.New())
	app.Use(middleware.Logger())
	app.Use(middlewares.SecurityHeaders(false))

	app.Use(func(c *fiber.Ctx) {
		if config.GetConfig().Server.EnableCSP == true {
			c.Set("Content-Security-Policy", "default-src 'self' https:; frame-ancestors 'none'; base-uri 'none'; form-action 'none';")
		}

		c.Next()
	})
}
