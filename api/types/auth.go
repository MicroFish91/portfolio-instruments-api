package types

import "github.com/gofiber/fiber/v3"

type AuthHandler interface {
	Login(fiber.Ctx) error
	Register(fiber.Ctx) error
}

// ---- Auth Response Types ----

type RegisterResponse struct {
	Data struct {
		User     User     `json:"user"`
		Settings Settings `json:"settings"`
		Status   string   `json:"status"`
	} `json:"data"`
	Error *string `json:"error"`
}

type LoginResponse struct {
	Data struct {
		Token  string `json:"token"`
		User   User   `json:"user"`
		Status string `json:"status"`
	} `json:"data"`
	Error *string `json:"error"`
}
