package holding

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GetHoldingsQuery struct {
	Holding_ids    []int               `json:"holding_ids"`
	Ticker         string              `json:"ticker"`
	Asset_category types.AssetCategory `json:"asset_category"`
	Is_deprecated  string              `json:"is_deprecated"`
}

func (q GetHoldingsQuery) Validate() error {
	return validation.ValidateStruct(&q,
		validation.Field(&q.Holding_ids),
		validation.Field(&q.Ticker, validation.Length(1, 32)),
		validation.Field(&q.Asset_category),
		validation.Field(&q.Is_deprecated, validation.In("true", "false")),
	)
}
