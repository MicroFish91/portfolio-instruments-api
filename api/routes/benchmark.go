package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func RegisterBenchmarkRoutes(r fiber.Router, h types.BenchmarkHandler) {
	r.Post("/benchmarks", h.CreateBenchmark, middleware.RequireAuth, middleware.AddBodyValidator[benchmark.CreateBenchmarkPayload]())
}
