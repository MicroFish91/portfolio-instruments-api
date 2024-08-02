package snapshot

import (
	"context"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/querybuilder"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

func (s *PostgresSnapshotStore) GetSnapshotTotal(ctx context.Context, userId, snapId int, options types.GetSnapshotTotalStoreOptions) (total float64, e error) {
	c, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	pgxb := querybuilder.NewPgxQueryBuilder()
	pgxb.AddQuery("SELECT SUM(total) AS snapshot_total")
	pgxb.AddQuery("FROM snapshots_values")
	pgxb.AddQueryWithPositionals("WHERE user_id = $x", []any{userId})
	pgxb.AddQueryWithPositionals("AND snap_id = $x", []any{snapId})

	if options.Omit_skip_reb {
		pgxb.AddQuery("AND skip_rebalance is false")
	}

	row := s.db.QueryRow(
		c,
		pgxb.Query,
		pgxb.QueryParams...,
	)

	var snapshot_total float64
	err := row.Scan(
		&snapshot_total,
	)

	if err != nil {
		return 0, err
	}

	return snapshot_total, nil
}
