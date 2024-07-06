package middleware

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

// Todo: Add support for verifying different roles
func RequireAuth(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	p, err := regexp.Compile(`^Bearer\s(\S+)`)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, fmt.Errorf("internal: %v", err.Error()))
	}

	matches := p.FindStringSubmatch(authHeader)
	if len(matches) < 2 {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("malformed authorization header"))
	}
	token := matches[1]

	jwtClaims, err := auth.VerifyJwt(token)
	if err != nil {
		return utils.SendError(c, fiber.StatusUnauthorized, err)
	}

	c.Locals(constants.LOCALS_REQ_USER, auth.AuthUserPayload{
		User_id: jwtClaims.UserId,
		Email:   jwtClaims.Email,
	})
	return c.Next()
}
