package types

import "github.com/gofiber/fiber/v3"

type AccountHandler interface {
	CreateAccount(fiber.Ctx) error
	GetAccounts(fiber.Ctx) error
	GetAccountById(fiber.Ctx) error
}
