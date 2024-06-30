package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type RegisterUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p RegisterUserPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Email, validation.Required, is.Email),
		validation.Field(&p.Password, validation.Required, validation.Length(5, 64)),
	)
}
