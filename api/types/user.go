package types

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type UserRole = string

const (
	Default UserRole = "DEFAULT"
	Admin   UserRole = "ADMIN"
)

type User struct {
	User_id        int       `json:"user_id,omitempty"`
	Email          string    `json:"email"`
	Enc_password   string    `json:"-"`
	User_role      UserRole  `json:"user_role"`
	Last_logged_in time.Time `json:"last_logged_in"`
	Verified       bool      `json:"verified"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

type UserHandler interface {
	GetUser(fiber.Ctx) error
	GetUsers(fiber.Ctx) error
	UpdateVerification(fiber.Ctx) error
	DeleteUser(fiber.Ctx) error
}

type UserStore interface {
	CreateUser(context.Context, *User) (User, error)
	GetUserById(ctx context.Context, userId int) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUsers(ctx context.Context, options *GetUsersStoreOptions) ([]User, PaginationMetadata, error)
	UpdateUserLoggedIn(ctx context.Context, userId int) (User, error)
	UpdateVerification(ctx context.Context, userId int) (User, error)
	DeleteUser(ctx context.Context, userId int) (User, error)
}

type GetUsersStoreOptions struct {
	Current_page int
	Page_size    int
}

// ---- User Response Types ----

type GetUserByIdResponse struct {
	Data struct {
		User User `json:"user"`
	} `json:"data"`
	Error string `json:"error"`
}

type DeleteUserResponse struct {
	Data struct {
		Message string `json:"message"`
		User    User   `json:"user"`
	} `json:"data"`
	Error string `json:"error"`
}
