package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DeleteAccountByIdParams struct {
	Id int `json:"id"`
}

func (p DeleteAccountByIdParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}
