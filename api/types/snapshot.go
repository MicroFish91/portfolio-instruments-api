package types

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Snapshot struct {
	Snap_id      int       `json:"snap_id"`
	Description  string    `json:"description,omitempty"`
	Snap_date    string    `json:"snap_date"`
	Total        float64   `json:"total"`
	Weighted_er  float64   `json:"weighted_er"`
	Benchmark_id int       `json:"benchmark_id,omitempty"`
	User_id      int       `json:"user_id"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type SnapshotValues struct {
	Snap_val_id    int       `json:"snap_val_id"`
	Snap_id        int       `json:"snap_id"`
	Account_id     int       `json:"account_id"`
	Holding_id     int       `json:"holding_id"`
	Total          float64   `json:"total"`
	Skip_rebalance bool      `json:"skip_rebalance"`
	User_id        int       `json:"user_id"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

type SnapshotHandler interface {
	GetSnapshots(fiber.Ctx) error
	GetSnapshotById(fiber.Ctx) error
	CreateSnapshot(fiber.Ctx) error
	RebalanceSnapshot(fiber.Ctx) error
}

type SnapshotStore interface {
	GetSnapshots(ctx context.Context, userId int, options *GetSnapshotsStoreOptions) (*[]Snapshot, *PaginationMetadata, error)
	GetSnapshotById(ctx context.Context, snapshotId, userId int) (*Snapshot, *[]SnapshotValues, error)
	CreateSnapshot(context.Context, *Snapshot) (*Snapshot, error)
	CreateSnapshotValues(context.Context, *SnapshotValues) (*SnapshotValues, error)
	GetSnapshotTotal(ctx context.Context, userId, snapId int, options GetSnapshotTotalStoreOptions) (total float64, err error)
	RefreshSnapshotTotal(ctx context.Context, userId, snapId int) (total float64, err error)
	RefreshSnapshotWeightedER(ctx context.Context, userId, snapId int) (weightedER float64, err error)
	TallyByAccount(ctx context.Context, userId, snapId int, options *GetTallyByAccountStoreOptions) (*ResourcesGrouped, error)
	TallyByHolding(ctx context.Context, userId, snapId int, options *GetTallyByHoldingStoreOptions) (*ResourcesGrouped, error)
}

type GetSnapshotsStoreOptions struct {
	Snap_ids        []int
	Snap_date_lower string
	Snap_date_upper string
	Order_date_by   string
	Current_page    int
	Page_size       int
}

type AccountsTallyCategory string

const (
	BY_ACCOUNT_NAME        AccountsTallyCategory = "ACCOUNT_NAME"
	BY_ACCOUNT_INSTITUTION AccountsTallyCategory = "ACCOUNT_INSTITUTION"
	BY_TAX_SHELTER         AccountsTallyCategory = "TAX_SHELTER"
)

type HoldingsTallyCategory string

const (
	BY_ASSET_CATEGORY HoldingsTallyCategory = "ASSET_CATEGORY"
)

type GetTallyByAccountStoreOptions struct {
	Tally_by AccountsTallyCategory
}

type GetTallyByHoldingStoreOptions struct {
	Tally_by HoldingsTallyCategory
	// Omit any snapshot_values that have "skip_rebalance" set to true
	Omit_skip_reb bool
}

type ResourcesGrouped struct {
	Fields []string  `json:"fields"`
	Total  []float64 `json:"total"`
}

type GetSnapshotTotalStoreOptions struct {
	Omit_skip_reb bool
}
