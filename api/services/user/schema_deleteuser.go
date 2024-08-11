package user

import validation "github.com/go-ozzo/ozzo-validation/v4"

type DeleteUserParams struct {
	Id int
}

func (p DeleteUserParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}
