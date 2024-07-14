package types

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

type Snapshot struct {
	Snap_id    int       `json:"snap_id"`
	Snap_date  string    `json:"snap_date"`
	Total      float64   `json:"total"`
	User_id    int       `json:"user_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
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
	CreateSnapshot(fiber.Ctx) error
}

type SnapshotStore interface {
	CreateSnapshot(*Snapshot) error
	CreateSnapshotValues(*SnapshotValues) error
	RefreshSnapshotTotal(snapshotId int) error
}
