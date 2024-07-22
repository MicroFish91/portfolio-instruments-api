package types

import (
	"context"
	"time"
)

type Settings struct {
	Settings_id    int       `json:"settings_id"`
	Reb_thresh_pct int       `json:"reb_thresh_pct"`
	Vp_thresh_pct  int       `json:"vp_thresh_pct,omitempty"`
	Vp_enabled     bool      `json:"vp_enabled"`
	User_id        int       `json:"user_id"`
	Benchmark_id   int       `json:"benchmark_id,omitempty"` // Int64 to match sql.NullInt64 type
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

type SettingsStore interface {
	CreateSettings(context.Context, *Settings) (*Settings, error)
}
