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
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger := logger.NewLogger(slog.LevelDebug) // Todo: Set log level via environment variable

	server := api.NewApiServer(config.Env.Port, db, logger)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
