package utils

import "github.com/gofiber/fiber/v3"

func SendJSON(c fiber.Ctx, status int, data any) error {
	c.Response().Header.SetContentType("application/json")
	return c.Status(status).JSON(fiber.Map{
		"status": status,
		"data":   data,
	})
}

func SendError(c fiber.Ctx, status int, err error) error {
	c.Response().Header.SetContentType("application/json")
	return c.Status(status).JSON(fiber.Map{
		"status": status,
		"error":  err.Error(),
	})
}
