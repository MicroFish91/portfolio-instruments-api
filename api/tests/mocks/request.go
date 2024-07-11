package mocks

import (
	"fmt"
	"net/http"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
)

func MockAuthRequestHeaders(r *http.Request) error {
	tok, err := auth.GenerateSignedJwt(1, "test@gmail.com")
	if err != nil {
		return err
	}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", tok))
	return nil
}
