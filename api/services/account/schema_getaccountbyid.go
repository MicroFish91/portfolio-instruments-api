package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type GetAccountByIdParams struct {
	Id int `json:"id"`
}

func (p GetAccountByIdParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}
