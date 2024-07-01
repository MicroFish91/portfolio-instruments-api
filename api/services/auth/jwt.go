package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET = "the-not-so-secret-placeholder-secret"

func GenerateSignedJwt(email string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "https://portfolioinstruments.com/api",
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
