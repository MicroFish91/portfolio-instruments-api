package benchmark

import (
	"github.com/jackc/pgx/v5"
)

type PostgresBenchmarkStore struct {
	db *pgx.Conn
}

func NewPostgresBenchmarkStore(db *pgx.Conn) *PostgresBenchmarkStore {
	return &PostgresBenchmarkStore{
		db: db,
	}
}
