package middleware

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func ParseAuthUserIfExists(jwtSecret string) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			// Only try to parse an auth token if it exists
			return c.Next()
		}

		p, err := regexp.Compile(`^Bearer\s(\S+)`)
		if err != nil {
			return utils.SendError(c, fiber.StatusInternalServerError, fmt.Errorf("internal: %v", err.Error()))
		}

		matches := p.FindStringSubmatch(authHeader)
		if len(matches) < 2 {
			return utils.SendError(c, fiber.StatusUnauthorized, errors.New("malformed authorization header"))
		}
		token := matches[1]

		jwtClaims, err := auth.VerifyJwt(token, jwtSecret)
		if err != nil {
			return utils.SendError(c, fiber.StatusUnauthorized, err)
		}

		c.Locals(constants.LOCALS_REQ_USER, auth.AuthUserPayload{
			User_id:   jwtClaims.UserId,
			Email:     jwtClaims.Email,
			User_role: jwtClaims.UserRole,
		})
		return c.Next()
	}
}

func RequireAuth(targetRole types.UserRole) func(fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		authUser, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
		if !ok || authUser.Email == "" || authUser.User_id == 0 {
			return utils.SendError(c, fiber.StatusUnauthorized, errors.New("an authenticated user is required to access this route"))
		}

		if getRolePriorityLevel(authUser.User_role) > getRolePriorityLevel(targetRole) {
			return utils.SendError(c, fiber.StatusForbidden, errors.New("authenticated user does not have sufficient privilege level to access this route"))
		}
		return c.Next()
	}
}

func getRolePriorityLevel(role types.UserRole) int {
	switch role {
	case types.Admin:
		return 1
	default:
		return 2
	}
}
