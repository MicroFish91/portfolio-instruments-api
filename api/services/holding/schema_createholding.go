package holding

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateHoldingPayload struct {
	Name           string              `json:"name"`
	Ticker         string              `json:"ticker"`
	Asset_category types.AssetCategory `json:"asset_category"`
	Expense_ratio  float32             `json:"expense_ratio"`
	Is_deprecated  bool                `json:"is_deprecated"`
}

func (p CreateHoldingPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(1, 256)),
		validation.Field(&p.Ticker, validation.Required, validation.Length(1, 32)),
		validation.Field(&p.Asset_category, validation.Required),
		validation.Field(&p.Expense_ratio, validation.Min(float32(0)), validation.Max(float32(100))),
		validation.Field(&p.Is_deprecated, validation.In(true, false)),
	)
}
