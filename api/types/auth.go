package types

import "github.com/gofiber/fiber/v3"

type AuthHandler interface {
	Login(fiber.Ctx) error
	Register(fiber.Ctx) error
}
