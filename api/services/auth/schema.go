package auth

type AuthUserPayload struct {
	User_id string `json:"user_id"`
	Email   string `json:"email"`
}
