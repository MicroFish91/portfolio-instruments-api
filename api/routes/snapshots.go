package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerSnapshotRoutes(r fiber.Router, h types.SnapshotHandler) {
	r.Get("/snapshots", h.GetSnapshots, middleware.RequireAuth)
	r.Post("/snapshots", h.CreateSnapshot, middleware.RequireAuth, middleware.AddBodyValidator[snapshot.CreateSnapshotPayload]())
}