package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerAuthRoutes(r fiber.Router, h types.AuthHandler) {
	r.Post("/login", h.Login, middleware.AddBodyValidator[auth.LoginPayload]())
	r.Post("/register", h.Register, middleware.AddBodyValidator[auth.RegisterPayload]())
}
