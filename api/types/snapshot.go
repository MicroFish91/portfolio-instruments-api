package types

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Snapshot struct {
	Snap_id     int     `json:"snap_id"`
	Description string  `json:"description,omitempty"`
	Snap_date   string  `json:"snap_date"`
	Total       float64 `json:"total"`
	// Todo: rename weighted_er_pct
	Weighted_er  float64   `json:"weighted_er"`
	Benchmark_id int       `json:"benchmark_id,omitempty"`
	User_id      int       `json:"user_id"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type SnapshotHandler interface {
	GetSnapshots(fiber.Ctx) error
	GetSnapshotById(fiber.Ctx) error
	CreateSnapshot(fiber.Ctx) error
	UpdateSnapshot(fiber.Ctx) error
	DeleteSnapshot(fiber.Ctx) error
	RebalanceSnapshot(fiber.Ctx) error
}

type SnapshotStore interface {
	GetSnapshots(ctx context.Context, userId int, options GetSnapshotsStoreOptions) ([]Snapshot, PaginationMetadata, error)
	GetSnapshotById(ctx context.Context, snapshotId, userId int) (Snapshot, []SnapshotValue, error)
	CreateSnapshot(context.Context, Snapshot) (Snapshot, error)
	UpdateSnapshot(context.Context, Snapshot) (Snapshot, error)
	DeleteSnapshot(ctx context.Context, snapshotId, userId int) (Snapshot, error)

	GetSnapshotTotal(ctx context.Context, userId, snapId int, options GetSnapshotTotalStoreOptions) (total float64, err error)
	RefreshSnapshotTotal(ctx context.Context, userId, snapId int) (total float64, err error)
	RefreshSnapshotWeightedER(ctx context.Context, userId, snapId int) (weightedER float64, err error)
	GroupByAccount(ctx context.Context, userId, snapId int, options GetGroupByAccountStoreOptions) (ResourcesGrouped, error)
	GroupByHolding(ctx context.Context, userId, snapId int, options GetGroupByHoldingStoreOptions) (ResourcesGrouped, error)
	GroupByMaturationDate(ctx context.Context, userId, snapId int, options GetGroupByMaturationDateStoreOptions) ([]MaturationDateResource, error)
}

type ResourcesGrouped struct {
	Fields []string  `json:"fields"`
	Total  []float64 `json:"total"`
}

type MaturationDateResource struct {
	Account_name    string  `json:"account_name"`
	Holding_name    string  `json:"holding_name"`
	Asset_category  string  `json:"asset_category"`
	Interest_rate   float64 `json:"interest_rate"`
	Maturation_date string  `json:"maturation_date"`
	Total           float64 `json:"total"`
	Skip_rebalance  bool    `json:"skip_rebalance"`
}

type GetSnapshotsStoreOptions struct {
	Snap_ids        []int
	Snap_date_lower string
	Snap_date_upper string
	Order_date_by   string
	Current_page    int
	Page_size       int
}

type AccountsGroupByCategory string

const (
	BY_ACCOUNT_NAME        AccountsGroupByCategory = "ACCOUNT_NAME"
	BY_ACCOUNT_INSTITUTION AccountsGroupByCategory = "ACCOUNT_INSTITUTION"
	BY_TAX_SHELTER         AccountsGroupByCategory = "TAX_SHELTER"
)

type HoldingsGroupByCategory string

const (
	BY_ASSET_CATEGORY HoldingsGroupByCategory = "ASSET_CATEGORY"
)

type GetGroupByAccountStoreOptions struct {
	Group_by AccountsGroupByCategory
}

type GetGroupByHoldingStoreOptions struct {
	Group_by HoldingsGroupByCategory
	// Omit any snapshots_values that have "skip_rebalance" set to true
	Omit_skip_reb bool
}

type GetGroupByMaturationDateStoreOptions struct {
	Maturation_start string
	Maturation_end   string
}

type GetSnapshotTotalStoreOptions struct {
	Omit_skip_reb bool
}
