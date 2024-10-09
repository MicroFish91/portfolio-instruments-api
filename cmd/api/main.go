package main

import (
	"log"
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api"
	"github.com/MicroFish91/portfolio-instruments-api/config"
	"github.com/MicroFish91/portfolio-instruments-api/db"
	"github.com/MicroFish91/portfolio-instruments-api/logger"
)

func main() {
	c := config.GetAppConfig()

	dbConfig := db.PostgresDbConfig{
		DbHost:           c.DbHost,
		DbPort:           c.DbPort,
		DbName:           c.DbName,
		DbUser:           c.DbUser,
		DbPassword:       c.DbPassword,
		DbSslMode:        c.DbSslMode,
		DbMaxConnections: c.DbMaxConnections,
		DbMinConnections: c.DbMinConnections,
	}

	db, err := db.NewPostgresStorage(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewLogger(slog.Level(c.LogLevel))
	apiConfig := &api.ApiConfig{
		Addr:                     c.Port,
		JwtSecret:                c.JwtSecret,
		RequireVerification:      c.RequireVerification,
		UnauthorizedRequestLimit: c.UnauthorizedRequestLimit,
		ShortRequestLimit:        c.ShortRequestLimit,
		LongRequestLimit:         c.LongRequestLimit,
	}

	server := api.NewApiServer(apiConfig, db, logger)
	defer server.Shutdown()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
