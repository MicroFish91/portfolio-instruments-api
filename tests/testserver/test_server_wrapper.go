package testserver

import (
	"context"
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api"
	"github.com/jackc/pgx/v5/pgxpool"
	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
)

type TestServerWrapper struct {
	tc         *pg.PostgresContainer
	TestServer *api.ApiServer
}

func newTestServerWrapper(addr string, db *pgxpool.Pool, logger *slog.Logger, tc *pg.PostgresContainer) *TestServerWrapper {
	return &TestServerWrapper{
		tc:         tc,
		TestServer: api.NewApiServer(addr, db, logger),
	}
}

func (w *TestServerWrapper) Shutdown() error {
	w.TestServer.Shutdown()
	if w.tc != nil {
		return w.tc.Terminate(context.Background())
	}
	return nil
}
