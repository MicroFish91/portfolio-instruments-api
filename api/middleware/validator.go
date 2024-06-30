package middleware

import (
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

type Validator interface {
	Validate() error
}

func AddRequestBodyValidator[T any]() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Parse request body
		var body T
		if err := utils.ParseRequestBody(c, &body); err != nil {
			return err
		}

		// Verify implementation of Validator interface
		v, ok := any(body).(Validator)
		if !ok {
			return fmt.Errorf("validation requires a implementation of the Validator interface")
		}

		// Validate
		if err := v.Validate(); err != nil {
			return err
		}

		c.Locals(constants.LOCALS_REQ_BODY, body)
		return c.Next()
	}
}
