package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerUserRoutes(r fiber.Router, h types.UserHandler) {
	// User
	r.Get("/me", h.GetMe, middleware.RequireAuth(types.Default))
	r.Get("/users", h.GetUsers, middleware.RequireAuth(types.Admin), middleware.AddQueryValidator[user.GetUsersQuery]())
	r.Get("/users/:id", h.GetUser, middleware.RequireAuth(types.Default), middleware.AddParamsValidator[user.GetUserParams]())
	r.Put("/users/:id/verify", h.UpdateVerification, middleware.RequireAuth(types.Admin), middleware.AddParamsValidator[user.UpdateVerificationParams]())
	r.Delete("/users/:id", h.DeleteUser, middleware.RequireAuth(types.Default), middleware.AddParamsValidator[user.DeleteUserParams]())

	// Settings
	r.Get("/users/:id/settings", h.GetSettings, middleware.RequireAuth(types.Default), middleware.AddParamsValidator[user.GetSettingsParams]())
	r.Put("/users/:id/settings", h.UpdateSettings, middleware.RequireAuth(types.Default), middleware.AddBodyValidator[user.UpdateSettingsPayload](), middleware.AddParamsValidator[user.UpdateSettingsParams]())
}
