package utils

import (
	"errors"
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/gofiber/fiber/v3"
)

func GetContextLogger(c fiber.Ctx) (*slog.Logger, error) {
	logger, ok := c.Locals(constants.LOCALS_LOGGER).(*slog.Logger)
	if !ok {
		return nil, errors.New("internal error: could not parse logger from context")
	}
	return logger, nil
}

func GetContextRequestId(c fiber.Ctx) (string, error) {
	id, ok := c.Locals(constants.LOCALS_REQ_ID).(string)
	if !ok {
		return "", errors.New("internal error: could not parse request id from context")
	}
	return id, nil
}
