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

type TestRegisterOptions struct {
	Parallel bool
}

func TestRegister(t *testing.T, p auth.RegisterPayload, options TestRegisterOptions) {
	if options.Parallel {
		t.Parallel()
	}

	tsw := testserver.GetTestServerWrapper()

	payload, err := json.Marshal(p)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/api/v1/register", bytes.NewBuffer(payload))
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

	var registerResponse types.RegisterResponse
	err = json.Unmarshal(bodyBytes, &registerResponse)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, res.StatusCode, fiber.StatusCreated)
	assert.Equal(t, p.Email, registerResponse.Data.User.Email)
}
