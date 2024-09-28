package testserver

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

var testServerWrapper *TestServerWrapper

func GetTestServerWrapper() *TestServerWrapper {
	if testServerWrapper == nil {
		testServerWrapper = initTestServerWrapper()
	}
	return testServerWrapper
}

func initTestServerWrapper() *TestServerWrapper {
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

	// Database connection (pgx driver)
	db, err := db.NewPostgresStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Server
	logger := logger.NewLogger(slog.LevelError)
	apiConfig := api.ApiConfig{
		Addr:              connStr,
		ShortRequestLimit: 99999,
		LongRequestLimit:  99999,
	}

	return newTestServerWrapper(apiConfig, db, logger, pgc)
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
