package snapshotvalue

import validation "github.com/go-ozzo/ozzo-validation"

type GetSnapshotValuesParams struct {
	Snap_id int `json:"snap_id"`
}

func (p GetSnapshotValuesParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Snap_id, validation.Required, validation.Min(1)),
	)
}
