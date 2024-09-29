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
	dbConfig := db.PostgresDbConfig{
		DbHost:           config.Env.DbHost,
		DbPort:           config.Env.DbPort,
		DbName:           config.Env.DbName,
		DbUser:           config.Env.DbUser,
		DbPassword:       config.Env.DbPassword,
		DbMaxConnections: config.Env.DbMaxConnections,
		DbMinConnections: config.Env.DbMinConnections,
	}

	db, err := db.NewPostgresStorage(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewLogger(slog.Level(config.Env.LogLevel))
	apiConfig := api.ApiConfig{
		Addr:                     config.Env.Port,
		JwtSecret:                config.Env.JwtSecret,
		UnauthorizedRequestLimit: config.Env.UnauthorizedRequestLimit,
		ShortRequestLimit:        config.Env.ShortRequestLimit,
		LongRequestLimit:         config.Env.LongRequestLimit,
	}

	server := api.NewApiServer(apiConfig, db, logger)
	defer server.Shutdown()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
