package routes

import (
	"github.com/gofiber/fiber"
	"github.com/lukewhrit/lynnx/database"
)

type createInput struct {
	Long string `json:"long" xml:"long" form:"long" query:"long"`
}

// Register contains all endpoints for the app
func Register(app *fiber.App) {
	api := app.Group("/v1")

	/*
	 * API "Schema"
	 * - GET /v1/:short
	 * - POST /v1/
	 */

	api.Get("/", func(c *fiber.Ctx) {
		// Just send back some information about the API
		c.JSON(&fiber.Map{
			"name":              "lynnx",
			"description":       "a powerful, accessible, fast and lightweight URL shortener.",
			"version":           "0.0.0",
			"project_url":       "https://github.com/lukewhrit/lynnx/",
			"documentation_url": "https://github.com/lukewhrit/lynnx/blob/master/spec.yml",
			"author":            "Luke Whrit <lukewhrit@gmail.com> (https://lukewhrit.xyz)",
		})
	})

	api.Get("/:short", func(c *fiber.Ctx) {
		c.SendStatus(501)
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

		if body.Long != "" {
			key, err := database.Create(body.Long)

			if err != nil {
				c.Status(500).JSON(&fiber.Map{
					"success": false,
					"error":   err.Error(),
				})
			}

			c.Status(201).JSON(&fiber.Map{
				"success": true,
				"short":   key,
			})
		} else {
			c.Status(400).JSON(&fiber.Map{
				"success": false,
				"error":   "\"long\" content field missing or empty.",
			})
		}
	})
}
