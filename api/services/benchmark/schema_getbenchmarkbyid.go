package benchmark

import validation "github.com/go-ozzo/ozzo-validation/v4"

type GetBenchmarkByIdParams struct {
	Id int
}

func (p GetBenchmarkByIdParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}
