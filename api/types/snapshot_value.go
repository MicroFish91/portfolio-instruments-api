package types

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type SnapshotValue struct {
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

type SnapshotValueHandler interface {
	CreateSnapshotValue(fiber.Ctx) error
	GetSnapshotValues(fiber.Ctx) error
	GetSnapshotValue(fiber.Ctx) error
	UpdateSnapshotValue(fiber.Ctx) error
	DeleteSnapshotValue(fiber.Ctx) error
}

type SnapshotValueStore interface {
	CreateSnapshotValue(context.Context, SnapshotValue) (SnapshotValue, error)
	GetSnapshotValues(ctx context.Context, snapId, userId int) ([]SnapshotValue, error)
	GetSnapshotValue(ctx context.Context, snapId, snapValId, userId int) (SnapshotValue, error)
	UpdateSnapshotValue(context.Context, SnapshotValue) (SnapshotValue, error)
	DeleteSnapshotValue(ctx context.Context, snapId, snapValId, userId int) (SnapshotValue, error)
}

// ---- Snapshot Value Response Types ----
type UpdateSnapshotValueResponse struct {
	Data struct {
		Snapshot_value      SnapshotValue `json:"snapshot_value"`
		Snapshot_total      float64       `json:"snapshot_total"`
		Snapshot_weighteder float64       `json:"snapshot_weighteder"`
	} `json:"data"`
	Error string `json:"error"`
}

type DeleteSnapshotValueResponse struct {
	Data struct {
		Snapshot_value SnapshotValue `json:"snapshot_value"`
		Message        string        `json:"message"`
	} `json:"data"`
	Error string `json:"error"`
}
