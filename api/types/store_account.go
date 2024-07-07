package types

type GetAccountsStoreOptions struct {
	AccountIds    []int
	TaxShelter    TaxShelter
	Institution   string
	Is_deprecated string
}

type AccountStore interface {
	CreateAccount(*Account) error
	GetAccounts(userId int, options GetAccountsStoreOptions) (*[]Account, error)
	GetAccountById(userId int, accountId int) (*Account, error)
}
