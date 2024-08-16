package tests

import (
	"os"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration"
	"github.com/MicroFish91/portfolio-instruments-api/tests/testserver"
)

func TestMain(m *testing.M) {
	tsw := testserver.GetTestServerWrapper()
	defer tsw.Shutdown()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestStart(t *testing.T) {
	// Ping
	t.Run("Ping", TestPing)

	// Integration tests (run in parallel)
	t.Run("User_AuthIntegration", integration.TestUser)
}
