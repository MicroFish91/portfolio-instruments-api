package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerSnapshotRoutes(r fiber.Router, h types.SnapshotHandler) {
	r.Get("/snapshots", h.GetSnapshots, middleware.RequireAuth, middleware.AddQueryValidator[snapshot.GetSnapshotsQuery]())
	r.Get("/snapshots/:id", h.GetSnapshotById, middleware.RequireAuth, middleware.AddParamsValidator[snapshot.GetSnapshotByIdParams](), middleware.AddQueryValidator[snapshot.GetSnapshotByIdQuery]())
	r.Get("/snapshots/:id/rebalance", h.RebalanceSnapshot, middleware.RequireAuth, middleware.AddParamsValidator[snapshot.GetSnapshotByIdParams]())
	r.Post("/snapshots", h.CreateSnapshot, middleware.RequireAuth, middleware.AddBodyValidator[snapshot.CreateSnapshotPayload]())
	r.Put("/snapshots/:id", h.UpdateSnapshot, middleware.RequireAuth, middleware.AddBodyValidator[snapshot.UpdateSnapshotPayload](), middleware.AddParamsValidator[snapshot.UpdateSnapshotParams]())
	r.Delete("/snapshots/:id", h.DeleteSnapshot, middleware.RequireAuth, middleware.AddParamsValidator[snapshot.DeleteSnapshotParams]())

	// r.Get("/snapshots/:sid/values")
	// r.Get("/snapshots/:sid/values/:svid")
	// r.Post("/snapshots/:sid/values/:svid")
	// r.Put("/snapshots/:sid/values/:svid")
	// r.Delete("/snapshots/:sid/values/:svid")
}
