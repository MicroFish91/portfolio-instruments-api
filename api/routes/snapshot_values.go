package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerSnapshotValueRoutes(r fiber.Router, h types.SnapshotValueHandler) {
	r.Post("/snapshots/:snap_id/values/", h.CreateSnapshotValue, middleware.RequireAuth, middleware.AddBodyValidator[snapshotvalue.CreateSnapshotValuePayload](), middleware.AddParamsValidator[snapshotvalue.CreateSnapshotValueParams]())
	r.Delete("/snapshots/:snap_id/values/:snap_val_id", h.DeleteSnapshotValue, middleware.RequireAuth, middleware.AddParamsValidator[snapshotvalue.DeleteSnapshotValueParams]())

	// r.Get("/snapshots/:sid/values")
	// r.Get("/snapshots/:sid/values/:svid")
	// r.Put("/snapshots/:sid/values/:svid")
}
