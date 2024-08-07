package benchmark

import validation "github.com/go-ozzo/ozzo-validation"

type DeleteBenchmarkById struct {
	Id int
}

func (p DeleteBenchmarkById) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}
