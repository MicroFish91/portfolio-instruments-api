package api

import (
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
	app := fiber.New()

	app.Get("/ping", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{"data": "pong"})
	})

	// apiv1 := app.Group("/api/v1")

	return app.Listen(s.addr)
}
