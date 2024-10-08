package benchmark

import (
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UpdateBenchmarkById struct {
	Id int
}

func (p UpdateBenchmarkById) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, validation.Min(1)),
	)
}

type UpdateBenchmarkPayload struct {
	Name             string                     `json:"name"`
	Description      string                     `json:"description"`
	Asset_allocation []types.AssetAllocationPct `json:"asset_allocation"`
	Std_dev_pct      float32                    `json:"std_dev_pct"`
	Real_return_pct  float32                    `json:"real_return_pct"`
	Drawdown_yrs     int                        `json:"drawdown_yrs"`
	Is_deprecated    bool                       `json:"is_deprecated"`
}

func (p UpdateBenchmarkPayload) Validate() error {
	var sum int
	for _, allocation := range p.Asset_allocation {
		err := validation.ValidateStruct(&allocation,
			validation.Field(&allocation.Category, validation.Required, validation.In(types.ValidAssetCategories...).Error("use a recognized asset category in all caps")),
			validation.Field(&allocation.Percent, validation.Required, validation.Min(1).Error("asset allocation percent must be a whole number greater than 0")),
		)

		if err != nil {
			return err
		}

		sum += allocation.Percent
	}

	if sum != 100 {
		return fmt.Errorf("asset allocation must sum to 100 but was %d", sum)
	}

	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required, validation.Length(1, 64)),
		validation.Field(&p.Description),
		validation.Field(&p.Std_dev_pct, validation.Min(float32(0)), validation.Max(float32(100))),
		validation.Field(&p.Real_return_pct, validation.Min(float32(0)), validation.Max(float32(100))),
		validation.Field(&p.Drawdown_yrs, validation.Min(0), validation.Max(50)),
		validation.Field(&p.Is_deprecated),
	)
}
