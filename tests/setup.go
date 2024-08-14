package tests

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api"
	"github.com/MicroFish91/portfolio-instruments-api/db"
	"github.com/MicroFish91/portfolio-instruments-api/logger"
	"github.com/MicroFish91/portfolio-instruments-api/migrator"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	pg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var ApiServer *api.ApiServer

func GetApiServer() *api.ApiServer {
	if ApiServer == nil {
		fmt.Println("found existing server")
		ApiServer = initApiServer()
	}
	return ApiServer
}

func initApiServer() *api.ApiServer {
	ctx := context.Background()

	// Postgres test container
	pgc, err := pg.Run(ctx,
		"postgres:16-alpine",
		pg.WithDatabase("pidb"),
		pg.WithUsername("piuser"),
		pg.WithPassword("pipassword"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(time.Second*5),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Connection string
	connStr, err := pgc.ConnectionString(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Migrations
	runDatabaseMigrations(connStr)

	// Database connection (pgx)
	db, err := db.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Server
	logger := logger.NewLogger(slog.LevelDebug)
	return api.NewApiServer(connStr, db, logger, pgc)

	// t.Cleanup(func() {
	// 	fmt.Println("----cleaning up postgres test container...----")
	// 	if err := pgc.Terminate(ctx); err != nil {
	// 		fmt.Printf("failed to terminate postgres test container: %s", err.Error())
	// 	}
	// })
}

func runDatabaseMigrations(connStr string) {
	m, err := migrator.NewPostgresMigrator(
		fmt.Sprintf("%ssslmode=disable", connStr), // connection string
		"file://../migrations",                    // migration source folder
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}