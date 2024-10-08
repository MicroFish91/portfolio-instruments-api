package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerAccountRoutes(r fiber.Router, h types.AccountHandler) {
	r.Get("/accounts", h.GetAccounts, middleware.RequireAuth(types.Default), middleware.AddQueryValidator[account.GetAccountsQuery]())
	r.Get("/accounts/:id", h.GetAccountById, middleware.RequireAuth(types.Default), middleware.AddParamsValidator[account.GetAccountByIdParams]())
	r.Post("/accounts", h.CreateAccount, middleware.RequireAuth(types.Default), middleware.AddBodyValidator[account.CreateAccountPayload]())
	r.Put("/accounts/:id", h.UpdateAccount, middleware.RequireAuth(types.Default), middleware.AddBodyValidator[account.UpdateAccountPayload](), middleware.AddParamsValidator[account.UpdateAccountByIdParams]())
	r.Delete("/accounts/:id", h.DeleteAccount, middleware.RequireAuth(types.Default), middleware.AddParamsValidator[account.DeleteAccountByIdParams]())
}
