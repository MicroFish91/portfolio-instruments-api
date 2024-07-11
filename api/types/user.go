package types

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

type User struct {
	User_id      int       `json:"user_id,omitempty"`
	Email        string    `json:"email"`
	Enc_password string    `json:"-"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type UserHandler interface {
	LoginUser(fiber.Ctx) error
	RegisterUser(fiber.Ctx) error
}

type UserStore interface {
	RegisterUser(*User) error
	GetUserByEmail(email string) (*User, error)
}
