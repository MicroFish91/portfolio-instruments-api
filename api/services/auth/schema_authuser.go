package auth

import "github.com/MicroFish91/portfolio-instruments-api/api/types"

type AuthUserPayload struct {
	User_id   int            `json:"user_id"`
	Email     string         `json:"email"`
	User_role types.UserRole `json:"user_role"`
}
