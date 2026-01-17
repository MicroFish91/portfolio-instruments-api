package snapshot

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdateSnapshotPayload struct {
	Snap_date               string `json:"snap_date"`
	Description             string `json:"description"`
	Rebalance_threshold_pct int    `json:"rebalance_threshold_pct"`
	Value_order             []int  `json:"value_order"`
	Benchmark_id            int    `json:"benchmark_id"`
}

type UpdateSnapshotParams struct {
	Id int `json:"id"`
}

func (p UpdateSnapshotPayload) Validate() error {
	if !regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`).Match([]byte(p.Snap_date)) {
		return errors.New("snap_date must follow string format mm/dd/yyyy")
	}

	return validation.ValidateStruct(&p,
		validation.Field(&p.Snap_date, validation.Length(10, 10)),
		validation.Field(&p.Description, validation.Length(1, 1024)),
		validation.Field(&p.Rebalance_threshold_pct, validation.Min(0), validation.Max(100)),
		validation.Field(&p.Value_order, validation.Length(1, 0)),
		validation.Field(&p.Benchmark_id, validation.Min(1)),
	)
}

func (p UpdateSnapshotParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Min(1)),
	)
}
