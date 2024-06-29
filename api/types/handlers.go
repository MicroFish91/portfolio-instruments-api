package types

import "github.com/gofiber/fiber/v3"

type UserHandler interface {
	HandleRegisterUser(fiber.Ctx) error
}
