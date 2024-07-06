package types

type UserStore interface {
	RegisterUser(*User) error
	GetUserByEmail(email string) (*User, error)
}

type GetAccountsStoreOptions struct {
	AccountIds  []int
	TaxShelter  TaxShelter
	Institution string
	Is_closed   string
}

type AccountStore interface {
	CreateAccount(*Account) error
	GetAccounts(userId int, options GetAccountsStoreOptions) (*[]Account, error)
	GetAccountById(userId int, accountId int) (*Account, error)
}
