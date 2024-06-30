package types

type UserStore interface {
	CreateUser(*User) error
	GetUserByEmail(email string) (*User, error)
}
