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
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

type Settings struct {
	Settings_id    int       `json:"settings_id"`
	Reb_thresh_pct int       `json:"reb_thresh_pct"`
	User_id        int       `json:"user_id"`
	Benchmark_id   int       `json:"benchmark_id,omitempty"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

type UserHandler interface {
	GetMe(fiber.Ctx) error
	GetUser(fiber.Ctx) error
	GetUsers(fiber.Ctx) error
	DeleteUser(fiber.Ctx) error

	GetSettings(fiber.Ctx) error
	UpdateSettings(fiber.Ctx) error
}

type UserStore interface {
	CreateUser(context.Context, User) (User, error)
	GetUserById(ctx context.Context, userId int) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUsers(ctx context.Context, options GetUsersStoreOptions) ([]User, PaginationMetadata, error)
	UpdateUserLoggedIn(ctx context.Context, userId int) (User, error)
	DeleteUser(ctx context.Context, userId int) (User, error)

	CreateSettings(context.Context, Settings) (Settings, error)
	GetSettings(ctx context.Context, userId int) (Settings, error)
	UpdateSettings(context.Context, Settings) (Settings, error)
}

type GetUsersStoreOptions struct {
	Current_page int
	Page_size    int
}

type GetMeResponse struct {
	Data struct {
		User     User     `json:"user"`
		Settings Settings `json:"settings"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetUserByIdResponse struct {
	Data struct {
		User     User     `json:"user"`
		Settings Settings `json:"settings"`
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

type GetSettingsResponse struct {
	Data struct {
		Settings Settings `json:"settings"`
	} `json:"data"`
	Error string `json:"error"`
}

type UpdateSettingsResponse struct {
	Data struct {
		Settings Settings `json:"settings"`
	} `json:"data"`
	Error string `json:"error"`
}
