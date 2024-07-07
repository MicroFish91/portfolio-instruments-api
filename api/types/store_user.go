package types

type UserStore interface {
	RegisterUser(*User) error
	GetUserByEmail(email string) (*User, error)
}
