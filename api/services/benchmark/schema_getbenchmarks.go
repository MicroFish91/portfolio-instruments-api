package benchmark

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type GetBenchmarksQuery struct {
	Benchmark_ids []int  `json:"benchmark_ids"`
	Name          string `json:"name"`
	Is_deprecated string `json:"is_deprecated"`

	PaginationQuery
}

type PaginationQuery struct {
	Current_page int `json:"current_page"`
	Page_size    int `json:"page_size"`
}

func (p PaginationQuery) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Current_page, validation.Min(1)),
		validation.Field(&p.Page_size, validation.Min(1)),
	)
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
