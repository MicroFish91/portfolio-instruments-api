package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func RegisterHoldingRoutes(r fiber.Router, h types.HoldingHandler) {
	r.Post("/holdings", h.CreateHolding, middleware.RequireAuth, middleware.AddBodyValidator[holding.CreateHoldingPayload]())
}
