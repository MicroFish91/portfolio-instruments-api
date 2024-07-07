package holding

import validation "github.com/go-ozzo/ozzo-validation/v4"

type GetHoldingByIdParams struct {
	Id int `json:"id"`
}

func (p GetHoldingByIdParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}
