package snapshot

import validation "github.com/go-ozzo/ozzo-validation/v4"

type GetSnapshotByIdParams struct {
	Id int `json:"id"`
}

func (p GetSnapshotByIdParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Min(1)),
	)
}
