package account

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
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

func (s *PostgresAccountStore) CreateAccount(a *types.Account) error {
	_, err := s.db.Exec(
		context.Background(),
		`INSERT INTO accounts
		(name, description, tax_shelter, institution, user_id)
		VALUES ($1, $2, $3, $4, $5)`,
		a.Name, a.Description, a.Tax_Shelter, a.Institution, a.User_id,
	)
	if err != nil {
		return err
	}
	return nil
}
