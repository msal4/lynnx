package routes

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/lukewhrit/lynnx/config"
	"github.com/lukewhrit/lynnx/database"
)

// Register contains all endpoints for the app
func Register(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/", func(c *fiber.Ctx) {
		// Just send back some information about the API
		c.JSON(&fiber.Map{
			"name":              "lynnx",
			"description":       "a powerful, accessible, fast and lightweight URL shortener.",
			"version":           "0.1.0",
			"project_url":       "https://github.com/lukewhrit/lynnx/",
			"documentation_url": "https://github.com/lukewhrit/lynnx/blob/master/spec.yml",
			"author":            "Luke Whrit <lukewhrit@gmail.com> (https://lukewhrit.xyz)",
		})
	})

	api.Get("/:short", func(c *fiber.Ctx) {
		// Make sure `short` is not empty and is of correct length.
		if c.Params("short") != "" && len(c.Params("short")) == config.GetConfig().LinkLength {
			value, err := database.Read(c.Params("short"))

			if err != nil {
				c.Status(500).JSON(&fiber.Map{
					"success": false,
					"error":   err.Error(),
				})

				return
			}

			c.Status(200).JSON(&fiber.Map{
				"success": true,
				"long":    value,
			})
		} else {
			c.Status(400).JSON(&fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("\"short\" parameter is missing, empty or of the wrong length (%d).", config.GetConfig().LinkLength),
			})

			return
		}
	})

	api.Post("/", func(c *fiber.Ctx) {
		body := new(createInput)

		if err := c.BodyParser(body); err != nil {
			c.Status(400).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})

			return
		}

		if err := body.validate(); err != nil {
			c.Status(400).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})

			return
		}

		key, err := database.Create(body.Long)

		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})

			return
		}

		c.Status(201).JSON(&fiber.Map{
			"success": true,
			"short":   key,
		})
	})
}
