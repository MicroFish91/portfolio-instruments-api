package types

import "time"

type User struct {
	User_id      string    `json:"user_id"`
	Email        string    `json:"email"`
	Enc_password string    `json:"-"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}
