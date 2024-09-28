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
	db, err := db.NewPostgresStorage("postgresql://localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewLogger(slog.LevelDebug) // Todo: Set log level via environment variable
	apiConfig := api.ApiConfig{
		Addr:              config.Env.Port,
		ShortRequestLimit: config.Env.ShortRequestLimit,
		LongRequestLimit:  config.Env.LongRequestLimit,
	}

	server := api.NewApiServer(apiConfig, db, logger)
	defer server.Shutdown()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
