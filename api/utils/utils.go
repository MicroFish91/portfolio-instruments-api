package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func ParseRequestBody(c fiber.Ctx, targetPayload interface{}) error {
	b := c.Body()
	if b == nil {
		return fmt.Errorf("missing request body")
	}

	return json.Unmarshal(b, targetPayload)
}
