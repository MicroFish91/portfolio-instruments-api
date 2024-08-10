package types

import validation "github.com/go-ozzo/ozzo-validation"

type PaginationMetadata struct {
	Current_page int `json:"current_page"`
	Page_size    int `json:"page_size"`
	Total_items  int `json:"total_items"`
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
