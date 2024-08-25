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

func TestApi(t *testing.T) {
	t.Run("Ping", TestPing)

	t.Run("Integration", func(t2 *testing.T) {
		t.Run("Users-Auth", integration.TestUserService)
		t.Run("Accounts", integration.TestAccountService)
	})
}
