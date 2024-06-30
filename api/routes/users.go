package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerUserRoutes(r fiber.Router, h types.UserHandler) {
	r.Post("/register", h.HandleRegisterUser, middleware.AddRequestBodyValidator[user.RegisterUserPayload]())
}
