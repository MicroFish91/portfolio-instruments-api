package main

import (
	"context"
	"log"

	"github.com/MicroFish91/portfolio-instruments-api/api"
	"github.com/MicroFish91/portfolio-instruments-api/db"
)

func main() {
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(context.Background())

	server := api.NewApiServer(":3000", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
