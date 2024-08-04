package holding

import validation "github.com/go-ozzo/ozzo-validation"

type DeleteHoldingParams struct {
	Id int `json:"id"`
}

func (p DeleteHoldingParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}
