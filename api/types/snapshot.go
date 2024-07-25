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
}

type SnapshotStore interface {
	GetSnapshots(ctx context.Context, userId int, options *GetSnapshotsStoreOptions) (*[]Snapshot, *PaginationMetadata, error)
	GetSnapshotById(ctx context.Context, snapshotId, userId int) (*Snapshot, *[]SnapshotValues, error)
	CreateSnapshot(context.Context, *Snapshot) (*Snapshot, error)
	CreateSnapshotValues(context.Context, *SnapshotValues) (*SnapshotValues, error)
	RefreshSnapshotTotal(ctx context.Context, userId, snapId int) (float64, error)
	TallyByAccount(ctx context.Context, userId, snapId int, options *GetTallyByAccountStoreOptions) (*AccountsGrouped, error)
}

type GetSnapshotsStoreOptions struct {
	Snap_ids        []int
	Snap_date_lower string
	Snap_date_upper string
	Order_date_by   string
	Current_page    int
	Page_size       int
}

type TallyCategory string

const (
	BY_ACCOUNT_NAME        TallyCategory = "ACCOUNT_NAME"
	BY_ACCOUNT_INSTITUTION TallyCategory = "ACCOUNT_INSTITUTION"
)

type GetTallyByAccountStoreOptions struct {
	Tally_by TallyCategory
}

type AccountsGrouped struct {
	Fields []string  `json:"fields"`
	Total  []float64 `json:"total"`
}
