package routes

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type createInput struct {
	Long string `json:"long" form:"long"`
}

func (c createInput) validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Long, validation.Required, is.URL),
	)
}
