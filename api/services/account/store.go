package account

import (
	"github.com/jackc/pgx/v5"
)

type PostgresAccountStore struct {
	db *pgx.Conn // Todo swap with pgxpool
}

func NewPostgresAccountStore(postgresDb *pgx.Conn) *PostgresAccountStore {
	return &PostgresAccountStore{
		db: postgresDb,
	}
}
