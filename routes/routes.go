package routes

import (
	"github.com/gofiber/fiber"
)

type createInput struct {
	Long string `json:"long" form:"long"`
}

// Register contains all endpoints for the app
func Register(app *fiber.App) {
	api := app.Group("/api/v1")

	// Register routes
	RegisterShorten(api)
	RegisterUnshorten(api)

	app.Get("/api", func(c *fiber.Ctx) {
		// Just send back some information about the API
		c.JSON(&fiber.Map{
			"name":              "lynnx",
			"description":       "a powerful, accessible, fast and lightweight URL shortener.",
			"version":           "1.0.0",
			"project_url":       "https://github.com/lukewhrit/lynnx/",
			"documentation_url": "https://github.com/lukewhrit/lynnx/blob/master/spec.yml",
			"author":            "Luke Whrit <lukewhrit@gmail.com> (https://lukewhrit.xyz)",
		})
	})
}
