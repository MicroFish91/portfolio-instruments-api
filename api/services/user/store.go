package user

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
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

func (s *PostgresUserStore) CreateUser(u *types.User) error {
	_, err := s.db.Exec(context.Background(), "INSERT INTO users (email, enc_password) VALUES ($1, $2)", u.Email, u.Enc_password)
	if err != nil {
		return err
	}
	return nil
}
