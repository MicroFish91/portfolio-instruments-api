package api

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/middleware"
	"github.com/MicroFish91/portfolio-instruments-api/api/routes"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
)

type ApiServer struct {
	addr string
	db   *pgx.Conn
}

func NewApiServer(addr string, db *pgx.Conn) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	fconfig := fiber.Config{
		ErrorHandler: middleware.FallbackHandler,
	}
	app := fiber.New(fconfig)

	// Initialize stores
	userStore := user.NewPostgresUserStore(s.db)
	accountStore := account.NewPostgresAccountStore(s.db)
	holdingStore := holding.NewPostgresHoldingStore(s.db)

	// Initialize handlers
	userHandler := user.NewUserHandler(userStore)
	accountHandler := account.NewAccountHandler(accountStore)
	holdingHandler := holding.NewHoldingHandler(holdingStore)

	routes.RegisterRoutes(app, userHandler, accountHandler, holdingHandler)
	return app.Listen(s.addr)
}
