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
		(name, description, tax_shelter, institution, is_closed, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		a.Name, a.Description, a.Tax_shelter, a.Institution, a.Is_closed, a.User_id,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresAccountStore) GetAccounts(userId int) (*[]types.Account, error) {
	rows, err := s.db.Query(
		context.Background(),
		`SELECT *
		FROM accounts
		WHERE user_id = $1`,
		userId,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var accounts []types.Account
	for rows.Next() {
		var a types.Account
		if err := rows.Scan(&a.Account_id, &a.Name, &a.Description, &a.Tax_shelter, &a.Institution, &a.Is_closed, &a.User_id, &a.Created_at, &a.Updated_at); err != nil {
			return nil, err
		}

		accounts = append(accounts, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &accounts, nil
}
