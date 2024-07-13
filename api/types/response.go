package types

type PaginationMetadata struct {
	Current_page int `json:"current_page"`
	Page_size    int `json:"page_size"`
	Total_items  int `json:"total_items"`
}
