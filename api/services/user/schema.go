package user

type RegisterUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
