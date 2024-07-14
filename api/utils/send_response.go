package utils

import (
	"errors"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/gofiber/fiber/v3"
)

func SendJSON(c fiber.Ctx, status int, data any) error {
	c.Response().Header.SetContentType("application/json")
	err := c.Status(status).JSON(fiber.Map{
		"status": status,
		"data":   data,
	})

	if err != nil {
		return err
	}
	return logOutboundTraffic(c)
}

func SendError(c fiber.Ctx, status int, e error) error {
	c.Response().Header.SetContentType("application/json")
	err := c.Status(status).JSON(fiber.Map{
		"status": status,
		"error":  e.Error(),
	})

	if err != nil {
		return err
	}
	return logOutboundTraffic(c)
}

func logOutboundTraffic(c fiber.Ctx) error {
	logger, err := GetContextLogger(c)
	if err != nil {
		return err
	}

	id, err := GetContextRequestId(c)
	if err != nil {
		return err
	}

	start, ok := c.Locals(constants.LOCALS_REQ_START).(time.Time)
	if !ok {
		return errors.New("internal error: unable to parse start time from context")
	}

	logger.Info(
		"Outbound traffic: ",
		"Id", id,
		"Method", c.Method(),
		"Path", c.Path(),
		"Ip", c.IP(),
		"Status", c.Response().StatusCode(),
		"Duration", time.Since(start),
	)

	return nil
}
