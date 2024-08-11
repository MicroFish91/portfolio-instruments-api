package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerUserRoutes(r fiber.Router, h types.UserHandler) {
	// User
	r.Get("/users/me", h.GetMe, middleware.RequireAuth)
	r.Delete("/users/:id", h.DeleteUser, middleware.RequireAuth, middleware.AddParamsValidator[user.DeleteUserParams]())

	// Settings
	r.Get("/users/settings", h.GetSettings, middleware.RequireAuth)
	r.Get("/users/:id/settings", h.GetSettings, middleware.RequireAuth)

	r.Put("/users/settings", h.UpdateSettings, middleware.RequireAuth, middleware.AddBodyValidator[user.UpdateSettingsPayload]())
	r.Put("/users/:id/settings", h.UpdateSettings, middleware.RequireAuth, middleware.AddBodyValidator[user.UpdateSettingsPayload]())
}
