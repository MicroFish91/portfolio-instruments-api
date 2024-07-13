package account

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type GetAccountsQuery struct {
	Ids           []int            `json:"ids"`
	Tax_shelter   types.TaxShelter `json:"tax_shelter"`
	Institution   string           `json:"institution"`
	Is_deprecated string           `json:"is_deprecated"`

	types.PaginationQuery
}

func (q GetAccountsQuery) Validate() error {
	err := q.PaginationQuery.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&q,
		validation.Field(&q.Tax_shelter),
		validation.Field(&q.Institution, validation.Length(1, 64)),
		validation.Field(&q.Is_deprecated, validation.In("true", "false")),
	)
}
