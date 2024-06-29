package types

type User struct {
	User_id      string `json:"user_id"`
	Email        string `json:"email"`
	Enc_password string `json:"enc_password"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}
