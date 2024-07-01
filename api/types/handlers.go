package types

import "github.com/gofiber/fiber/v3"

type UserHandler interface {
	HandleLoginUser(fiber.Ctx) error
	HandleRegisterUser(fiber.Ctx) error
}
