package api

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/routes"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApiServer struct {
	addr   string
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewApiServer(addr string, db *pgxpool.Pool, logger *slog.Logger) *ApiServer {
	return &ApiServer{
		addr:   addr,
		db:     db,
		logger: logger,
	}
}

func GetFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: middleware.FallbackHandler,
	}
}

func (s *ApiServer) Run() error {
	app := fiber.New(GetFiberConfig())

	// Middleware
	app.Use(middleware.AddIncomingTrafficLogger(s.logger))
	app.Use(middleware.AddLocalsContextLogger(s.logger))

	// Stores
	userStore := user.NewPostgresUserStore(s.db, s.logger)
	accountStore := account.NewPostgresAccountStore(s.db, s.logger)
	holdingStore := holding.NewPostgresHoldingStore(s.db, s.logger)
	benchmarkStore := benchmark.NewPostgresBenchmarkStore(s.db, s.logger)
	snapshotStore := snapshot.NewPostgresSnapshotStore(s.db, s.logger)
	snapshotValueStore := snapshotvalue.NewPostgresSnapshotValueStore(s.db, s.logger)

	// Handlers
	userHandler := user.NewUserHandler(userStore, benchmarkStore, s.logger)
	accountHandler := account.NewAccountHandler(accountStore, s.logger)
	holdingHandler := holding.NewHoldingHandler(holdingStore, s.logger)
	benchmarkHandler := benchmark.NewBenchmarkHandler(userStore, benchmarkStore, s.logger)
	snapshotHandler := snapshot.NewSnapshotHandler(userStore, benchmarkStore, accountStore, holdingStore, snapshotStore, snapshotValueStore, s.logger)
	snapshotValueHandler := snapshotvalue.NewSnapshotValueHandler(accountStore, holdingStore, snapshotStore, snapshotValueStore, s.logger)

	// Routes
	routes.RegisterRoutes(app, userHandler, accountHandler, holdingHandler, benchmarkHandler, snapshotHandler, snapshotValueHandler)
	return app.Listen(s.addr)
}
