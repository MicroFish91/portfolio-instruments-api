package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresStorage() (*pgxpool.Pool, error) {
	connStr := "postgresql://localhost:5432/postgres"

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal(err)
	}

	config.MaxConns = 5

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	return pgxpool.NewWithConfig(ctx, config)
}
