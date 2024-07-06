package middleware

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func AddBodyValidator[T any]() fiber.Handler {
	return func(c fiber.Ctx) error {
		var body T
		if err := parseRequestBody(c, &body); err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, fmt.Errorf("error parsing request payload: %v", err.Error()))
		}

		v, ok := any(body).(Validator)
		if !ok {
			return utils.SendError(c, fiber.StatusInternalServerError, errors.New("internal: request payload validation requires an implementation of the Validator interface"))
		}

		if err := v.Validate(); err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, fmt.Errorf("invalid request payload: %v", err.Error()))
		}

		c.Locals(constants.LOCALS_REQ_BODY, body)
		return c.Next()
	}
}

func parseRequestBody(c fiber.Ctx, targetPayload interface{}) error {
	b := c.Body()
	if b == nil {
		return fmt.Errorf("missing request body")
	}
	return json.Unmarshal(b, targetPayload)
}
