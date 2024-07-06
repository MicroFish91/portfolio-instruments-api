package types

type UserStore interface {
	CreateUser(*User) error
	GetUserByEmail(email string) (*User, error)
}

type AccountStore interface {
	CreateAccount(*Account) error
	GetAccounts(userId int, accountIds []int) (*[]Account, error)
	GetAccountById(userId int, accountId int) (*Account, error)
}
