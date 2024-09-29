package api

import (
	"log/slog"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/routes"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ApiConfig struct {
	Addr                     string
	UnauthorizedRequestLimit int
	ShortRequestLimit        int
	LongRequestLimit         int
}

type ApiServer struct {
	cfg    ApiConfig
	db     *pgxpool.Pool
	logger *slog.Logger
	App    *fiber.App
}

func NewApiServer(cfg ApiConfig, db *pgxpool.Pool, logger *slog.Logger) *ApiServer {
	api := &ApiServer{
		cfg:    cfg,
		db:     db,
		logger: logger,
	}
	api.init()
	return api
}

func getFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: middleware.FallbackHandler,
	}
}

func (s *ApiServer) init() {
	s.App = fiber.New(getFiberConfig())

	// Middleware
	s.App.Use(middleware.AddIncomingTrafficLogger(s.logger))
	s.App.Use(middleware.AddLocalsContextLogger(s.logger))
	s.App.Use(middleware.ParseAuthUserIfExists)
	s.App.Use(middleware.AddUnauthorizedRateLimiter(s.cfg.UnauthorizedRequestLimit, 60*time.Minute))
	s.App.Use(middleware.AddRateLimiter(s.cfg.ShortRequestLimit, 1*time.Minute))
	s.App.Use(middleware.AddRateLimiter(s.cfg.LongRequestLimit, 30*time.Minute))

	// Stores
	userStore := user.NewPostgresUserStore(s.db, s.logger)
	accountStore := account.NewPostgresAccountStore(s.db, s.logger)
	holdingStore := holding.NewPostgresHoldingStore(s.db, s.logger)
	benchmarkStore := benchmark.NewPostgresBenchmarkStore(s.db, s.logger)
	snapshotStore := snapshot.NewPostgresSnapshotStore(s.db, s.logger)
	snapshotValueStore := snapshotvalue.NewPostgresSnapshotValueStore(s.db, s.logger)

	// Handlers
	authHandler := auth.NewAuthHandler(userStore, s.logger)
	userHandler := user.NewUserHandler(userStore, benchmarkStore, s.logger)
	accountHandler := account.NewAccountHandler(accountStore, s.logger)
	holdingHandler := holding.NewHoldingHandler(holdingStore, s.logger)
	benchmarkHandler := benchmark.NewBenchmarkHandler(userStore, benchmarkStore, s.logger)
	snapshotHandler := snapshot.NewSnapshotHandler(userStore, benchmarkStore, accountStore, holdingStore, snapshotStore, snapshotValueStore, s.logger)
	snapshotValueHandler := snapshotvalue.NewSnapshotValueHandler(accountStore, holdingStore, snapshotStore, snapshotValueStore, s.logger)

	// Routes
	routes.RegisterRoutes(s.App, authHandler, userHandler, accountHandler, holdingHandler, benchmarkHandler, snapshotHandler, snapshotValueHandler)
}

func (s *ApiServer) Run() error {
	return s.App.Listen(s.cfg.Addr)
}

func (s *ApiServer) Shutdown() {
	s.db.Close()
}
