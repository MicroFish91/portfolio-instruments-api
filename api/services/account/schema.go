package account

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type CreateAccountPayload struct {
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Shelter_type types.TaxShelter `json:"shelter_type"`
	Institution  string           `json:"institution"`
}

func (p CreateAccountPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(1, 64)),
		validation.Field(&p.Description, validation.Length(1, 1024)),
		validation.Field(&p.Shelter_type, validation.Required, validation.In(types.TAXABLE, types.TRADITIONAL, types.ROTH, types.HSA, types.FIVE_TWENTY_NINE)),
		validation.Field(&p.Institution, validation.Required, validation.Length(1, 64)),
	)
}
