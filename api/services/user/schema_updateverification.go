package user

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UpdateVerificationParams struct {
	Id int
}

func (p UpdateVerificationParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}
