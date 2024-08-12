package user

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *UserHandlerImpl) GetUsers(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user request body"))
	}

	queryPayload, ok := c.Locals(constants.LOCALS_REQ_QUERY).(GetUsersQuery)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid query params from request"))
	}

	if userPayload.User_role != types.Admin {
		return utils.SendError(c, fiber.StatusForbidden, errors.New("insufficient permissions level to access this route"))
	}

	users, pagination, err := h.userStore.GetUsers(c.Context(), types.GetUsersStoreOptions{
		Current_page: queryPayload.Current_page,
		Page_size:    queryPayload.Page_size,
	})
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"users":      users,
		"pagination": pagination,
	})
}