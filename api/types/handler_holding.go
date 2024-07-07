package types

import "github.com/gofiber/fiber/v3"

type HoldingHandler interface {
	CreateHolding(fiber.Ctx) error
	GetHoldings(fiber.Ctx) error
}
