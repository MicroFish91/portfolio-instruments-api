package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func registerSnapshotRoutes(r fiber.Router, h types.SnapshotHandler) {
	r.Get("/snapshots", h.GetSnapshots, middleware.RequireAuth(types.Default), middleware.AddQueryValidator[snapshot.GetSnapshotsQuery]())
	r.Get("/snapshots/:id", h.GetSnapshotById, middleware.RequireAuth(types.Default), middleware.AddParamsValidator[snapshot.GetSnapshotByIdParams](), middleware.AddQueryValidator[snapshot.GetSnapshotByIdQuery]())
	r.Get("/snapshots/:id/rebalance", h.RebalanceSnapshot, middleware.RequireAuth(types.Default), middleware.AddParamsValidator[snapshot.GetSnapshotByIdParams]())
	r.Post("/snapshots", h.CreateSnapshot, middleware.RequireAuth(types.Default), middleware.AddBodyValidator[snapshot.CreateSnapshotPayload]())
	r.Put("/snapshots/:id", h.UpdateSnapshot, middleware.RequireAuth(types.Default), middleware.AddBodyValidator[snapshot.UpdateSnapshotPayload](), middleware.AddParamsValidator[snapshot.UpdateSnapshotParams]())
	r.Put("/snapshots/:id/order", h.UpdateValueOrder, middleware.RequireAuth(types.Default), middleware.AddBodyValidator[snapshot.UpdateValueOrderPayload](), middleware.AddParamsValidator[snapshot.UpdateValueOrderParams]())
	r.Delete("/snapshots/:id", h.DeleteSnapshot, middleware.RequireAuth(types.Default), middleware.AddParamsValidator[snapshot.DeleteSnapshotParams]())
}
