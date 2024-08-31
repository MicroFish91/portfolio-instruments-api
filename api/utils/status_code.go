package utils

import (
	"regexp"

	"github.com/gofiber/fiber/v3"
)

func StatusCodeFromError(err error) int {
	switch {
	case regexp.MustCompile(`(?i)bad_request`).Match([]byte(err.Error())):
		return fiber.StatusBadRequest
	case regexp.MustCompile(`(?i)no[\s]rows[\s]in[\s]result[\s]set`).Match([]byte(err.Error())):
		return fiber.StatusNotFound
	case regexp.MustCompile(`(?i)duplicate`).Match([]byte(err.Error())) ||
		regexp.MustCompile(`(?i)violates[\s]*unique`).Match([]byte(err.Error())):
		return fiber.StatusConflict
	case regexp.MustCompile(`(?i)deadline[\s]*exceeded`).Match([]byte(err.Error())):
		return fiber.StatusGatewayTimeout
	default:
		return fiber.StatusInternalServerError
	}
}
