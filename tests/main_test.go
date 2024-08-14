package tests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	s := GetApiServer()
	defer s.Shutdown()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestStart(t *testing.T) {
	s := GetApiServer()
	fmt.Println(s)

	req, err := http.NewRequest(http.MethodGet, "/ping", bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Fatal(err)
	}

	res, err := s.App.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println("Response Body:", bodyString)
}
