package user

import (
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
	// Todo add check for existing user

	err := h.store.CreateUser(&types.User{
		Email:        userPayload.Email,
		Enc_password: userPayload.Password, // Todo Needs encryption
	})

	if err != nil {
		// Todo return a properly formatted error
		return err
	}

	// Todo return a properly formatted response
	return c.Status(fiber.StatusCreated).JSON(nil)
}
