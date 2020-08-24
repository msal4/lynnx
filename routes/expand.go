package routes

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/lukewhrit/lynnx/config"
	"github.com/lukewhrit/lynnx/database"
)

func expand(c *fiber.Ctx) {
	// Make sure `short` is not empty and is of correct length.
	if c.Params("short") != "" && len(c.Params("short")) == config.Configuration.LinkLength {
		value, err := database.GetLink(c.Params("short"))

		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err.Error(),
			})

			return
		}

		c.Redirect(value)
	} else {
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("\"short\" parameter is missing, empty or of the wrong length (%d).", config.Configuration.LinkLength),
		})

		return
	}
}
