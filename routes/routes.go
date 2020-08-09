package routes

import (
	"github.com/gofiber/fiber"
)

// Register contains all endpoints for the app
func Register(app *fiber.App) {
	api := app.Group("/api")

	// Register routes
	RegisterShorten(api)
	RegisterExpand(api)

	api.Get("/", func(c *fiber.Ctx) {
		// Just send back some information about the API
		c.JSON(&fiber.Map{
			"name":              "lynnx",
			"description":       "a powerful, accessible, fast and lightweight URL shortener.",
			"version":           "0.1.0",
			"project_url":       "https://github.com/lukewhrit/lynnx/",
			"documentation_url": "https://github.com/lukewhrit/lynnx/blob/main/spec.yml",
			"author":            "Luke Whrit <lukewhrit@gmail.com> (https://lukewhrit.xyz)",
		})
	})
}
