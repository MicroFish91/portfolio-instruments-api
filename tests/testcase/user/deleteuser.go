package user

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUser(t *testing.T, route string, token string, userId int, expectedStatusCode int) {
	var deleteUserResponse types.DeleteUserResponse
	var r string

	if route == "" {
		r = fmt.Sprintf("/api/v1/users/%d", userId)
	} else {
		r = route
	}

	res := utils.SendDeleteRequest(t, r, token, &deleteUserResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, res.StatusCode, fiber.StatusOK)
		assert.NotEmpty(t, deleteUserResponse.Data.Message)
		assert.EqualExportedValues(
			t,
			types.User{
				User_id:        userId,
				Email:          deleteUserResponse.Data.User.Email,
				User_role:      types.Default,
				Last_logged_in: deleteUserResponse.Data.User.Last_logged_in,
				Created_at:     deleteUserResponse.Data.User.Created_at,
				Updated_at:     deleteUserResponse.Data.User.Updated_at,
			},
			deleteUserResponse.Data.User,
		)
	case 400:
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	case 401:
		assert.Equal(t, fiber.StatusUnauthorized, res.StatusCode)
	case 403:
		assert.Equal(t, fiber.StatusForbidden, res.StatusCode)
	case 404:
		assert.Equal(t, fiber.StatusNotFound, res.StatusCode)
	default:
		t.Fatal("provided an unexpected status code")
	}
}
