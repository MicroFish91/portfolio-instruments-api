package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerUserRoutes(r fiber.Router, h types.UserHandler) {
	// GET /me or /users/:id
	// DEL /users/:id
	// PUT & PATCH /users/:id
	r.Post("/login", h.HandleLoginUser, middleware.AddRequestBodyValidator[user.LoginUserPayload]())
	r.Post("/register", h.HandleRegisterUser, middleware.AddRequestBodyValidator[user.RegisterUserPayload]())
}
