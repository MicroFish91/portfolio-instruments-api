package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/testserver"
)

// params
// p = incoming payload
// body = unmarshalled response body
func SendAuthRequest(t *testing.T, route string, p any, body any) *http.Response {
	tsw := testserver.GetTestServerWrapper()

	payload, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, route, bytes.NewBuffer(payload))
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

	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		t.Fatal(err)
	}

	return res
}

func SendCreateOrUpdateRequest(t *testing.T, method string, route string, token string, p any, body any) *http.Response {
	if method != http.MethodPost && method != http.MethodPut {
		t.Fatal("invalid send request method")
	}

	tsw := testserver.GetTestServerWrapper()

	payload, err := json.Marshal(p)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(method, route, bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := tsw.TestServer.App.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		t.Fatal(err)
	}

	return res
}

func SendGetRequest(t *testing.T, route string, token string, body any) *http.Response {
	tsw := testserver.GetTestServerWrapper()

	req, err := http.NewRequest(http.MethodGet, route, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := tsw.TestServer.App.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bodyBytes, &body)

	if err != nil {
		t.Fatal(err)
	}
	return res
}

func SendDeleteRequest(t *testing.T, route string, token string, body any) *http.Response {
	tsw := testserver.GetTestServerWrapper()

	req, err := http.NewRequest(http.MethodDelete, route, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := tsw.TestServer.App.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(bodyBytes, &body)

	if err != nil {
		t.Fatal(err)
	}
	return res
}
