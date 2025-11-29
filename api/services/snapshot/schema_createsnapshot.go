package snapshot

import (
	"errors"
	"regexp"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateSnapshotPayload struct {
	Snap_date               string                                     `json:"snap_date"`
	Description             string                                     `json:"description"`
	Snapshot_values         []snapshotvalue.CreateSnapshotValuePayload `json:"snapshot_values"`
	Rebalance_threshold_pct int                                        `json:"rebalance_threshold_pct"`
	Benchmark_id            int                                        `json:"benchmark_id"`
}

func (p CreateSnapshotPayload) Validate() error {
	if len(p.Snapshot_values) == 0 {
		return errors.New("must provide a list of values for snapshot_values")
	}

	for _, snapshotValue := range p.Snapshot_values {
		if err := snapshotValue.Validate(); err != nil {
			return err
		}
	}

	if !regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`).Match([]byte(p.Snap_date)) {
		return errors.New("snap_date must follow string format mm/dd/yyyy")
	}

	snapDate, err := time.Parse("01/02/2006", p.Snap_date)
	if err != nil {
		return errors.New("snap_date must follow string format mm/dd/yyyy")
	}

	now := time.Now()
	if snapDate.After(now) {
		return errors.New("snap_date cannot be a date in the future")
	}

	return validation.ValidateStruct(&p,
		validation.Field(&p.Snap_date, validation.Length(10, 10)),
		validation.Field(&p.Description, validation.Length(1, 1024)),
		validation.Field(&p.Snapshot_values),
		validation.Field(&p.Rebalance_threshold_pct, validation.Min(0), validation.Max(100)),
		validation.Field(&p.Benchmark_id, validation.Min(1)),
	)
}
