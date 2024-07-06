package auth

type AuthUserPayload struct {
	User_id int    `json:"user_id"`
	Email   string `json:"email"`
}
