package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/testserver"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

type LoginTestOptions struct {
	Parallel bool
}

func TestLogin(t *testing.T, p auth.LoginPayload, options LoginTestOptions) {
	if options.Parallel {
		t.Parallel()
	}

	tsw := testserver.GetTestServerWrapper()

	payload, err := json.Marshal(p)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v1/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	res, err := tsw.TestServer.App.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var loginResponse types.LoginResponse
	err = json.Unmarshal(bodyBytes, &loginResponse)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, fiber.StatusCreated)
	assert.Equal(t, p.Email, loginResponse.Data.User.Email)
}
