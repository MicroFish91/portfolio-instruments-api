package account

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdateAccountPayload struct {
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Tax_shelter   types.TaxShelter `json:"tax_shelter"`
	Institution   string           `json:"institution"`
	Is_deprecated bool             `json:"is_deprecated"`
}

func (p UpdateAccountPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(1, 64)),
		validation.Field(&p.Description, validation.Length(1, 1024)),
		validation.Field(&p.Tax_shelter, validation.Required, validation.In(types.ValidTaxShelters...).Error("use a recognized tax shelter in all caps")),
		validation.Field(&p.Institution, validation.Required, validation.Length(1, 64)),
		validation.Field(&p.Is_deprecated, validation.In(true, false)),
	)
}

type UpdateAccountByIdParams struct {
	Id int `json:"id"`
}

func (p UpdateAccountByIdParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}
