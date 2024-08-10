package snapshotvalue

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DeleteSnapshotValueParams struct {
	Snap_id     int `json:"snap_id"`
	Snap_val_id int `json:"snap_val_id"`
}

func (p DeleteSnapshotValueParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Snap_id, validation.Required, validation.Min(1)),
		validation.Field(&p.Snap_val_id, validation.Required, validation.Min(1)),
	)
}
