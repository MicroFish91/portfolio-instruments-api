package user

import (
	"github.com/jackc/pgx/v5"
)

type PostgresUserStore struct {
	db *pgx.Conn // Todo swap with pgxpool
}

func NewPostgresUserStore(postgresDb *pgx.Conn) *PostgresUserStore {
	return &PostgresUserStore{
		db: postgresDb,
	}
}
