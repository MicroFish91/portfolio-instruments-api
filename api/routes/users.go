package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerAuthRoutes(r fiber.Router, h types.UserHandler) {
	r.Post("/login", h.HandleLoginUser, middleware.AddBodyValidator[user.LoginUserPayload]())
	r.Post("/register", h.HandleRegisterUser, middleware.AddBodyValidator[user.RegisterUserPayload]())
}
