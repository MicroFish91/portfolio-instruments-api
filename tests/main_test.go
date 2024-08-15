package tests

import (
	"os"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/testserver"
)

func TestMain(m *testing.M) {
	tsw := testserver.GetTestServerWrapper()
	defer tsw.Shutdown()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestStart(t *testing.T) {
	t.Run("Ping", TestPing)
	// Test Auth
	// Test Users
	// etc...
}
