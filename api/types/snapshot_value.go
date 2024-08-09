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

// todo: finish delete snapshot value and then migrate types for createsnapshotvalue

type SnapshotValueHandler interface {
	CreateSnapshotValue(fiber.Ctx) error
	DeleteSnapshotValue(fiber.Ctx) error
}

type SnapshotValueStore interface {
	CreateSnapshotValue(context.Context, SnapshotValue) (SnapshotValue, error)
	DeleteSnapshotValue(ctx context.Context, snapValId, userId int) (SnapshotValue, error)
}
