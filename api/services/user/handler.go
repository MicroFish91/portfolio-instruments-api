package user

import (
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

type UserHandlerImpl struct {
	store types.UserStore
}

func NewUserHandler(store types.UserStore) *UserHandlerImpl {
	return &UserHandlerImpl{
		store: store,
	}
}

func (h *UserHandlerImpl) HandleRegisterUser(c fiber.Ctx) error {
	var userPayload RegisterUserPayload
	if err := utils.ParseRequestBody(c, &userPayload); err != nil {
		// Todo return a properly formatted error
		return err
	}

	// Todo add request validation

	user, _ := h.store.GetUserByEmail(userPayload.Email)
	if user != nil {
		// Todo return a properly formatted error
		return fmt.Errorf("user with provided email already exists")
	}

	encPassword, err := auth.HashPassword(userPayload.Password)
	if err != nil {
		// Todo return a properly formatted error
		return err
	}

	err = h.store.CreateUser(&types.User{
		Email:        userPayload.Email,
		Enc_password: encPassword,
	})

	if err != nil {
		// Todo return a properly formatted error
		return err
	}

	// Todo return a properly formatted response
	return c.Status(fiber.StatusCreated).JSON(nil)
}
