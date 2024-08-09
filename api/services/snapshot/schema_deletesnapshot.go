package snapshot

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DeleteSnapshotParams struct {
	Id int `json:"id"`
}

func (p DeleteSnapshotParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Min(1)),
	)
}
