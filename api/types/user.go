package types

import (
	"context"
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

type Settings struct {
	Settings_id    int       `json:"settings_id"`
	Reb_thresh_pct int       `json:"reb_thresh_pct"`
	Vp_thresh_pct  int       `json:"vp_thresh_pct,omitempty"`
	Vp_enabled     bool      `json:"vp_enabled"`
	User_id        int       `json:"user_id"`
	Benchmark_id   int       `json:"benchmark_id,omitempty"`
	Created_at     time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

type UserHandler interface {
	LoginUser(fiber.Ctx) error
	RegisterUser(fiber.Ctx) error
	GetSettings(fiber.Ctx) error
}

type UserStore interface {
	RegisterUser(context.Context, *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateSettings(context.Context, *Settings) (*Settings, error)
	GetSettings(ctx context.Context, userId int) (*Settings, error)
}
