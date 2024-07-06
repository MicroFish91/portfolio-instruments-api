package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerAccountRoutes(r fiber.Router, h types.AccountHandler) {
	r.Get("/accounts", h.GetAccounts, middleware.AuthValidator, middleware.AddQueryValidator[account.GetAccountsQuery]())
	r.Get("/accounts/:id", h.GetAccountById, middleware.AuthValidator, middleware.AddParamsValidator[account.GetAccountByIdParams]())
	r.Post("/accounts", h.CreateAccount, middleware.AuthValidator, middleware.AddBodyValidator[account.CreateAccountPayload]())
}
