package tests

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/testserver"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	t.Parallel()
	tsw := testserver.GetTestServerWrapper()

	req, err := http.NewRequest(http.MethodGet, "/ping", bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Fatal(err)
	}

	res, err := tsw.TestServer.App.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assert.Equal(t, res.StatusCode, fiber.StatusOK)
}
