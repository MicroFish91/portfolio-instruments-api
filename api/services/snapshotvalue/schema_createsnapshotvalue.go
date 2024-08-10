package snapshotvalue

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateSnapshotValuePayload struct {
	Account_id     int     `json:"account_id"`
	Holding_id     int     `json:"holding_id"`
	Total          float64 `json:"total"`
	Skip_rebalance bool    `json:"skip_rebalance"`
}

func (p CreateSnapshotValuePayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Account_id, validation.Required),
		validation.Field(&p.Holding_id, validation.Required),
		validation.Field(&p.Total, validation.Required),
		validation.Field(&p.Skip_rebalance),
	)
}

type CreateSnapshotValueParams struct {
	Snap_id int `json:"snap_id"`
}

func (p CreateSnapshotValueParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Snap_id, validation.Required, validation.Min(1)),
	)
}
