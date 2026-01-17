package snapshot

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdateValueOrderPayload struct {
	Value_order []int `json:"value_order"`
}

type UpdateValueOrderParams struct {
	Id int `json:"id"`
}

func (p UpdateValueOrderPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Value_order, validation.Length(1, 0)),
	)
}

func (p UpdateValueOrderParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Min(1)),
	)
}
