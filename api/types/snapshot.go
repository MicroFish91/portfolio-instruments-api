package types

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Snapshot struct {
	Snap_id         int       `json:"snap_id"`
	Description     string    `json:"description,omitempty"`
	Snap_date       string    `json:"snap_date"`
	Total           float64   `json:"total"`
	Weighted_er_pct float64   `json:"weighted_er_pct"`
	Benchmark_id    int       `json:"benchmark_id,omitempty"`
	User_id         int       `json:"user_id"`
	Created_at      time.Time `json:"created_at"`
	Updated_at      time.Time `json:"updated_at"`
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
	GetSnapshots(ctx context.Context, userId int, options *GetSnapshotsStoreOptions) ([]Snapshot, PaginationMetadata, error)
	GetSnapshotById(ctx context.Context, snapshotId, userId int) (Snapshot, []SnapshotValue, error)
	CreateSnapshot(context.Context, *Snapshot) (Snapshot, error)
	UpdateSnapshot(context.Context, *Snapshot) (Snapshot, error)
	DeleteSnapshot(ctx context.Context, snapshotId, userId int) (Snapshot, error)

	GetSnapshotByDate(ctx context.Context, snapshotDate string, userId int) (Snapshot, error)
	GetSnapshotTotal(ctx context.Context, userId, snapId int, options *GetSnapshotTotalStoreOptions) (total float64, err error)
	RefreshSnapshotTotal(ctx context.Context, userId, snapId int) (total float64, err error)
	RefreshSnapshotWeightedER(ctx context.Context, userId, snapId int) (weightedER float64, err error)
	GroupByAccount(ctx context.Context, userId, snapId int, options *GetGroupByAccountStoreOptions) (ResourcesGrouped, error)
	GroupByHolding(ctx context.Context, userId, snapId int, options *GetGroupByHoldingStoreOptions) (ResourcesGrouped, error)
	GroupByMaturationDate(ctx context.Context, userId, snapId int, options *GetGroupByMaturationDateStoreOptions) ([]MaturationDateResource, error)
	GroupByLiquidity(ctx context.Context, userId, snapId int) (resources []LiquidityResource, total float64, err error)
}

type ResourcesGrouped struct {
	Fields []string  `json:"fields"`
	Total  []float64 `json:"total"`
}

type MaturationDateResource struct {
	Account_name      string  `json:"account_name"`
	Holding_name      string  `json:"holding_name"`
	Asset_category    string  `json:"asset_category"`
	Interest_rate_pct float64 `json:"interest_rate_pct"`
	Maturation_date   string  `json:"maturation_date"`
	Total             float64 `json:"total"`
	Skip_rebalance    bool    `json:"skip_rebalance"`
}

type LiquidityResource struct {
	Account_name   string  `json:"account_name"`
	Holding_name   string  `json:"holding_name"`
	Asset_category string  `json:"asset_category"`
	Ticker         string  `json:"ticker"`
	Institution    string  `json:"institution"`
	TaxShelter     string  `json:"tax_shelter"`
	Total          float64 `json:"total"`
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

type AssetAllocation struct {
	Category string  `json:"category"`
	Value    float64 `json:"value"`
}

// ---- Snapshot Response Types ----
type CreateSnapshotResponse struct {
	Data struct {
		Snapshot        Snapshot        `json:"snapshot"`
		Snapshot_values []SnapshotValue `json:"snapshot_values"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetSnapshotResponse struct {
	Data struct {
		Snapshot        Snapshot        `json:"snapshot"`
		Snapshot_values []SnapshotValue `json:"snapshot_values"`
		Accounts        []Account       `json:"accounts"`
		Holdings        []Holding       `json:"holdings"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetSnapshotsResponse struct {
	Data struct {
		Snapshots  []Snapshot         `json:"snapshots"`
		Pagination PaginationMetadata `json:"pagination"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetSnapshotAccountsResponse struct {
	Data struct {
		Accounts_grouped ResourcesGrouped `json:"accounts_grouped"`
		Field_type       string           `json:"field_type"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetSnapshotHoldingsResponse struct {
	Data struct {
		Holdings_grouped ResourcesGrouped `json:"holdings_grouped"`
		Field_type       string           `json:"field_type"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetSnapshotMaturationDateResponse struct {
	Data struct {
		Resources        []MaturationDateResource `json:"resources"`
		Field_type       string                   `json:"field_type"`
		Maturation_start string                   `json:"maturation_start"`
		Maturation_end   string                   `json:"maturation_end"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetSnapshotLiquidityResponse struct {
	Data struct {
		Resources    []LiquidityResource `json:"resources"`
		Liquid_total string              `json:"liquid_total"`
		Field_type   string              `json:"field_type"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetSnapshotRebalanceResponse struct {
	Data struct {
		Target_allocation         *[]AssetAllocation `json:"target_allocation"`
		Current_allocation        *[]AssetAllocation `json:"current_allocation"`
		Change_required           *[]AssetAllocation `json:"change_required"`
		Rebalance_thresh_pct      int                `json:"rebalance_thresh_pct"`
		Snapshot_total            float64            `json:"snapshot_total"`
		Snapshot_total_omit_skips float64            `json:"snapshot_total_omit_skips"`
	} `json:"data"`
	Error string `json:"error"`
}

type UpdateSnapshotResponse struct {
	Data struct {
		Snapshot Snapshot `json:"snapshot"`
	} `json:"data"`
	Error string `json:"error"`
}

type DeleteSnapshotResponse struct {
	Data struct {
		Message  string   `json:"message"`
		Snapshot Snapshot `json:"snapshot"`
	} `json:"data"`
	Error string `json:"error"`
}
