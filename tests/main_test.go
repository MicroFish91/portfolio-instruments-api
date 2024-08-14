package tests

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	s := GetApiServer()
	defer s.Shutdown()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestRun(t *testing.T) {
	s := GetApiServer()
	fmt.Println(s)
}
