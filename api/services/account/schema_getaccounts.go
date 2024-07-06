package account

type GetAccountsQuery struct {
	Ids []int `json:"ids"`
}

func (p GetAccountsQuery) Validate() error {
	return nil
}
