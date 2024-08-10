package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerSnapshotValueRoutes(r fiber.Router, h types.SnapshotValueHandler) {
	r.Post("/snapshots/:snap_id/values", h.CreateSnapshotValue, middleware.RequireAuth, middleware.AddBodyValidator[snapshotvalue.CreateSnapshotValuePayload](), middleware.AddParamsValidator[snapshotvalue.CreateSnapshotValueParams]())
	r.Get("/snapshots/:snap_id/values", h.GetSnapshotValues, middleware.RequireAuth, middleware.AddParamsValidator[snapshotvalue.GetSnapshotValuesParams]())
	r.Get("/snapshots/:snap_id/values/:snap_val_id", h.GetSnapshotValue, middleware.RequireAuth, middleware.AddParamsValidator[snapshotvalue.GetSnapshotValueParams]())
	r.Put("/snapshots/:snap_id/values/:snap_val_id", h.UpdateSnapshotValue, middleware.RequireAuth, middleware.AddBodyValidator[snapshotvalue.UpdateSnapshotValuePayload](), middleware.AddParamsValidator[snapshotvalue.UpdateSnapshotValueParams]())
	r.Delete("/snapshots/:snap_id/values/:snap_val_id", h.DeleteSnapshotValue, middleware.RequireAuth, middleware.AddParamsValidator[snapshotvalue.DeleteSnapshotValueParams]())
}
