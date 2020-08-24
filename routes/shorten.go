package routes

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/lukewhrit/lynnx/database"
	"github.com/lukewhrit/lynnx/utils"
)

func shorten(c *fiber.Ctx) {
	if c.FormValue("long") != "" && utils.IsURL(c.FormValue("long")) {
		key, err := database.NewLink(c.FormValue("long"))

		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})

			return
		}

		c.Status(201).JSON(&fiber.Map{
			"success": false,
			"long":    key,
		})
	} else {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("\"long\" parameter is missing, empty, or not a valid URL"),
		})
	}
}
