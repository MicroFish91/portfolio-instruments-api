package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerUserRoutes(r fiber.Router, h types.UserHandler) {
	// Auth
	r.Post("/login", h.LoginUser, middleware.AddBodyValidator[user.LoginUserPayload]())
	r.Post("/register", h.RegisterUser, middleware.AddBodyValidator[user.RegisterUserPayload]())

	// Settings
	r.Get("/users/settings", h.GetSettings, middleware.RequireAuth)
	r.Get("/users/:id/settings", h.GetSettings, middleware.RequireAuth)
}
