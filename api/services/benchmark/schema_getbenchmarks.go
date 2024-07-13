package benchmark

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	validation "github.com/go-ozzo/ozzo-validation"
)

type GetBenchmarksQuery struct {
	Benchmark_ids []int  `json:"benchmark_ids"`
	Name          string `json:"name"`
	Is_deprecated string `json:"is_deprecated"`

	types.PaginationQuery
}

func (q GetBenchmarksQuery) Validate() error {
	err := q.PaginationQuery.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&q,
		validation.Field(&q.Benchmark_ids),
		validation.Field(&q.Name, validation.Length(1, 64)),
		validation.Field(&q.Is_deprecated, validation.In("true", "false")),
	)
}
