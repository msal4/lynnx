package main

import (
	"fmt"
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/lukewhrit/lynnx/config"
	"github.com/lukewhrit/lynnx/middlewares"
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

	// Start the server
	address := fmt.Sprintf("%s:%x", config.GetConfig().Server.Host, config.GetConfig().Server.Port)

	log.Fatal(app.Listen(address))
}

func registerMiddleware(app *fiber.App) {
	app.Use(middleware.Compress(middleware.CompressConfig{
		Level: config.GetConfig().CompressionLevel,
	}))

	app.Use(cors.New())
	app.Use(middlewares.SecurityHeaders())
	app.Use(middleware.Logger())
}
