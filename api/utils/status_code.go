package utils

import (
	"regexp"

	"github.com/gofiber/fiber/v3"
)

func StatusCodeFromError(err error) int {
	switch {
	case regexp.MustCompile(`(?i)no[\s]rows[\s]in[\s]result[\s]set`).Match([]byte(err.Error())):
		return fiber.StatusNotFound
	case regexp.MustCompile(`(?i)deadline[\s]*exceeded`).Match([]byte(err.Error())):
		return fiber.StatusGatewayTimeout
	default:
		return fiber.StatusInternalServerError
	}
}
