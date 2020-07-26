package main

import (
	"fmt"
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/lukewhrit/lynnx/config"
	"github.com/lukewhrit/lynnx/middlewares"
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
	registerEndpoints(app)

	// Serve static files
	app.Static("/", "./static", fiber.Static{
		Compress: true,
	})

	// Start the server
	address := fmt.Sprintf("%s:%x", config.GetServer().Host, config.GetServer().Port)

	fmt.Println(address)

	log.Fatal(app.Listen(address))
}

func registerMiddleware(app *fiber.App) {
	app.Use(middleware.Compress(middleware.CompressConfig{
		Level: config.GetServer().CompressionLevel,
	}))

	app.Use(cors.New())
	app.Use(middlewares.SecurityHeaders())
	app.Use(middleware.Logger())
}

func registerEndpoints(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})
}
