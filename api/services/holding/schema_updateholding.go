package holding

import (
	"errors"
	"regexp"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type UpdateHoldingPayload struct {
	Name            string              `json:"name"`
	Ticker          string              `json:"ticker,omitempty"`
	Asset_category  types.AssetCategory `json:"asset_category"`
	Expense_ratio   float32             `json:"expense_ratio,omitempty"`
	Maturation_date string              `json:"maturation_date,omitempty"`
	Interest_rate   float32             `json:"interest_rate,omitempty"`
	Is_deprecated   bool                `json:"is_deprecated"`
}

func (p UpdateHoldingPayload) Validate() error {
	if p.Maturation_date != "" {
		pattern := regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`)
		ok := pattern.Match([]byte(p.Maturation_date))
		if !ok {
			return errors.New("maturation_date must follow string format mm/dd/yyyy")
		}
	}

	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(1, 256)),
		validation.Field(&p.Ticker, validation.Length(1, 32)),
		validation.Field(&p.Asset_category, validation.Required, validation.In(types.ValidAssetCategories...).Error("not a valid asset category")),
		validation.Field(&p.Expense_ratio, validation.Min(float32(0)), validation.Max(float32(100))),
		validation.Field(&p.Interest_rate, validation.Min(float32(0)), validation.Max(float32(100))),
		validation.Field(&p.Maturation_date, validation.Length(10, 10)),
		validation.Field(&p.Is_deprecated),
	)
}

type UpdateHoldingParams struct {
	Id int `json:"id"`
}

func (p UpdateHoldingParams) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}