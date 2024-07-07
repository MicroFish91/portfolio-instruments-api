package holding

import (
	"github.com/jackc/pgx/v5"
)

type PostgresHoldingStore struct {
	db *pgx.Conn
}

func NewPostgresHoldingStore(db *pgx.Conn) *PostgresHoldingStore {
	return &PostgresHoldingStore{
		db: db,
	}
}
