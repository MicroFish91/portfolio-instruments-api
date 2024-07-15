package snapshot

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreateSnapshotPayload struct {
	Snap_date       string                        `json:"snap_date"`
	Snapshot_values []CreateSnapshotValuesPayload `json:"snapshot_values"`
}

type CreateSnapshotValuesPayload struct {
	Account_id     int     `json:"account_id"`
	Holding_id     int     `json:"holding_id"`
	Total          float64 `json:"total"`
	Skip_rebalance bool    `json:"skip_rebalance"`
}

func (p CreateSnapshotPayload) Validate() error {
	for _, snapshotValue := range p.Snapshot_values {
		if err := snapshotValue.Validate(); err != nil {
			return err
		}
	}

	return validation.ValidateStruct(&p,
		validation.Field(&p.Snap_date, validation.Required),
		validation.Field(&p.Snapshot_values),
	)
}

func (p CreateSnapshotValuesPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Account_id, validation.Required),
		validation.Field(&p.Holding_id, validation.Required),
		validation.Field(&p.Total, validation.Required),
		validation.Field(&p.Skip_rebalance),
	)
}
