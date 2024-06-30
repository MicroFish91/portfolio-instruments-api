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

func (s *PostgresUserStore) GetUserByEmail(email string) (*types.User, error) {
	var u types.User

	row := s.db.QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1", email)
	if err := row.Scan(&u.User_id, &u.Email, &u.Enc_password, &u.Created_at, &u.Updated_at); err != nil {
		return nil, err
	}

	return &u, nil
}
