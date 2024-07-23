package user

import validation "github.com/go-ozzo/ozzo-validation/v4"

type UpdateSettingsPayload struct {
	Reb_thresh_pct int  `json:"reb_thresh_pct"`
	Vp_thresh_pct  int  `json:"vp_thresh_pct,omitempty"`
	Vp_enabled     bool `json:"vp_enabled"`
	Benchmark_id   int  `json:"benchmark_id,omitempty"`
}

func (p UpdateSettingsPayload) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Reb_thresh_pct, validation.Min(0), validation.Max(50)),
		validation.Field(&p.Vp_thresh_pct, validation.Min(0), validation.Max(50)),
		validation.Field(&p.Vp_enabled),
		validation.Field(&p.Benchmark_id, validation.Min(1)),
	)
}
