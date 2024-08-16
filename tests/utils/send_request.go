package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/testserver"
)

// params
// p = incoming payload
// body = unmarshalled response body
func SendPostRequest(t *testing.T, route string, p any, body any) *http.Response {
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
