package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_SECRET = "the-not-so-secret-jwt-secret"

type JwtClaims struct {
	UserId int
	Email  string
	jwt.RegisteredClaims
}

func GenerateSignedJwt(userId int, email string) (string, error) {
	claims := JwtClaims{
		UserId: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "https://portfolioinstruments.com/api",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func VerifyJwt(tokenString string) (*JwtClaims, error) {
	t, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(*JwtClaims)
	if !ok {
		return nil, errors.New("failed to parse jwt claims")
	}

	ok = claims.ExpiresAt.After(time.Now())
	if !ok {
		return nil, errors.New("the provided token has expired")
	}

	return claims, nil
}
