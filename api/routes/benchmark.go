package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func RegisterBenchmarkRoutes(r fiber.Router, h types.BenchmarkHandler) {
	r.Get("/benchmarks", h.GetBenchmarks, middleware.RequireAuth, middleware.AddQueryValidator[benchmark.GetBenchmarksQuery]())
	r.Get("/benchmarks/:id", h.GetBenchmarkById, middleware.RequireAuth, middleware.AddParamsValidator[benchmark.GetBenchmarkByIdParams]())
	r.Post("/benchmarks", h.CreateBenchmark, middleware.RequireAuth, middleware.AddBodyValidator[benchmark.CreateBenchmarkPayload]())
}
