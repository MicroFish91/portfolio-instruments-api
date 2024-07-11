package benchmark

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateBenchmarkPayload struct {
	Name             string                  `json:"name"`
	Description      string                  `json:"description"`
	Asset_allocation []types.AssetAllocation `json:"asset_allocation"`
	Std_dev_pct      float32                 `json:"std_dev_pct"`
	Real_return_pct  float32                 `json:"real_return_pct"`
	Drawdown_yrs     int                     `json:"drawdown_yrs"`
	Is_deprecated    bool                    `json:"is_deprecated"`
}

func (p CreateBenchmarkPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(1, 64)),
		validation.Field(&p.Description),
		validation.Field(&p.Asset_allocation, validation.Required),
		validation.Field(&p.Std_dev_pct, validation.Min(float32(0)), validation.Max(float32(100))),
		validation.Field(&p.Real_return_pct, validation.Min(float32(0)), validation.Max(float32(100))),
		validation.Field(&p.Drawdown_yrs, validation.Min(0), validation.Max(50)),
		validation.Field(&p.Is_deprecated, validation.In(true, false)),
	)
}
