package user

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
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

func (h *UserHandlerImpl) HandleLoginUser(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(LoginUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user request body"))
	}

	user, err := h.store.GetUserByEmail(userPayload.Email)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	if err := auth.CompareHashAndPassword(user.Enc_password, userPayload.Password); err != nil {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("invalid login credentials"))
	}

	jwt, err := auth.GenerateSignedJwt(user.Email)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"token": jwt,
		"user":  user,
	})
}

func (h *UserHandlerImpl) HandleRegisterUser(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(RegisterUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user request body"))
	}

	user, _ := h.store.GetUserByEmail(userPayload.Email)
	if user != nil {
		return utils.SendError(c, fiber.StatusConflict, errors.New("user with provided email already exists"))
	}

	encpw, err := auth.HashPassword(userPayload.Password)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	err = h.store.CreateUser(&types.User{
		Email:        userPayload.Email,
		Enc_password: encpw,
	})
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{"message": "user registered successfully"})
}
