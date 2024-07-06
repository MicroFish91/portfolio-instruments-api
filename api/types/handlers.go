package types

import "github.com/gofiber/fiber/v3"

type UserHandler interface {
	LoginUser(fiber.Ctx) error
	RegisterUser(fiber.Ctx) error
}

type AccountHandler interface {
	CreateAccount(fiber.Ctx) error
	GetAccounts(fiber.Ctx) error
	GetAccountById(fiber.Ctx) error
}
