package routes

import (
	"github.com/gofiber/fiber"
)

// Register contains all endpoints for the app
func Register(app *fiber.App) {
	app.Get("/:short", expand)
	app.Post("/", shorten)
}
