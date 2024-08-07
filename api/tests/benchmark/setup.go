package benchmark

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api"
	"github.com/MicroFish91/portfolio-instruments-api/api/routes"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/tests/mocks"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

var app *fiber.App

func registerAppWithBenchmarks() *fiber.App {
	if app != nil {
		return app
	}

	logger := mocks.NewMockLogger(slog.LevelDebug)
	app = fiber.New(api.GetFiberConfig())

	userStore := mocks.NewMockUserStore()
	benchmarkStore := mocks.NewMockBenchmarkStore()
	benchmarkHandler := benchmark.NewBenchmarkHandler(userStore, benchmarkStore, logger)

	registerMockBenchmarkRoutes(app, benchmarkHandler)
	return app
}

func registerMockBenchmarkRoutes(app *fiber.App, benchmarkHandler types.BenchmarkHandler) {
	v1 := app.Group("/api/v1")
	routes.RegisterBenchmarkRoutes(v1, benchmarkHandler)
}
