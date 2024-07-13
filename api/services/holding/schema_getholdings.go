package holding

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GetHoldingsQuery struct {
	Holding_ids              []int               `json:"holding_ids"`
	Ticker                   string              `json:"ticker"`
	Asset_category           types.AssetCategory `json:"asset_category"`
	Is_deprecated            string              `json:"is_deprecated"`
	Has_maturation_remaining string              `json:"has_maturation_remaining"`

	types.PaginationQuery
}

func (q GetHoldingsQuery) Validate() error {
	err := q.PaginationQuery.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&q,
		validation.Field(&q.Holding_ids),
		validation.Field(&q.Ticker, validation.Length(1, 32)),
		validation.Field(&q.Asset_category),
		validation.Field(&q.Is_deprecated, validation.In("true", "false")),
		validation.Field(&q.Has_maturation_remaining, validation.In("true", "false")),
	)
}
