package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewPostgresStorage() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), "postgresql://localhost:5432/postgres")
}
