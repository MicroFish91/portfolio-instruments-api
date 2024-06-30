package user

import (
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
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
	userPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(RegisterUserPayload)
	if !ok {
		return fmt.Errorf("unable to parse valid user request body")
	}

	user, _ := h.store.GetUserByEmail(userPayload.Email)
	if user != nil {
		// Todo return a properly formatted error
		return fmt.Errorf("user with provided email already exists")
	}

	encpw, err := auth.HashPassword(userPayload.Password)
	if err != nil {
		// Todo return a properly formatted error
		return err
	}

	err = h.store.CreateUser(&types.User{
		Email:        userPayload.Email,
		Enc_password: encpw,
	})

	if err != nil {
		// Todo return a properly formatted error
		return err
	}

	// Todo return a properly formatted response
	return c.Status(fiber.StatusCreated).JSON(nil)
}
