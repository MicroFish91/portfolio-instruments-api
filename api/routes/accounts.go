package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerAccountRoutes(r fiber.Router, h types.AccountHandler) {
	r.Get("/accounts", h.HandleGetAccounts, middleware.AuthValidator)
	r.Post("/accounts", h.HandleCreateAccount, middleware.AuthValidator, middleware.AddRequestBodyValidator[account.CreateAccountPayload]())
}
