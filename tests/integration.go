package tests

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/auth"
)

func TestIntegration(t *testing.T) {
	t.Run("Auth", auth.TestAuth)
	// Test Auth
	// Test Users
	// etc...
}
