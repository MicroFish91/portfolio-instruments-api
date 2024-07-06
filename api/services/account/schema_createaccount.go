package account

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateAccountPayload struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Tax_shelter types.TaxShelter `json:"tax_shelter"`
	Institution string           `json:"institution"`
	Is_closed   bool             `json:"is_closed"`
}

func (p CreateAccountPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(1, 64)),
		validation.Field(&p.Description, validation.Length(1, 1024)),
		validation.Field(&p.Tax_shelter, validation.Required, validation.In(types.TAXABLE, types.TRADITIONAL, types.ROTH, types.HSA, types.FIVE_TWENTY_NINE)),
		validation.Field(&p.Institution, validation.Required, validation.Length(1, 64)),
		validation.Field(&p.Is_closed, validation.In(true, false)),
	)
}
