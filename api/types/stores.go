package types

type UserStore interface {
	CreateUser(*User) error
}
