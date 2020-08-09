package routes

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofiber/fiber"
	"github.com/lukewhrit/lynnx/database"
)

type createInput struct {
	Long string `json:"long" form:"long"`
}

func (c createInput) validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Long, validation.Required, is.URL),
	)
}

// RegisterShorten contains an endpoint that allows for links to be shortened
func RegisterShorten(api fiber.Router) {
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

		key, err := database.NewLink(body.Long)

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
