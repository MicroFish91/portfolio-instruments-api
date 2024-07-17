package utils

import (
	"regexp"

	"github.com/gofiber/fiber/v3"
)

func StatusCodeFromError(err error) int {
	switch {
	case regexp.MustCompile(`deadline[\s]*exceeded`).Match([]byte(err.Error())):
		return fiber.StatusGatewayTimeout
	default:
		return fiber.StatusInternalServerError
	}
}
