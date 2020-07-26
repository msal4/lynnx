package routes

import "github.com/gofiber/fiber"

// Register contains all endpoints for the app
func Register(app *fiber.App) {
	api := app.Group("/v1")

	api.Get("/", func(c *fiber.Ctx) {
		c.JSON(fiber.Map{
			"message": "Hello, World 👋!",
		})
	})
}
