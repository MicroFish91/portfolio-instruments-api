package auth

import (
	"errors"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserId   int
	Email    string
	UserRole types.UserRole
	jwt.RegisteredClaims
}

func GenerateSignedJwt(userId int, email string, role types.UserRole, jwtSecret string) (string, error) {
	claims := JwtClaims{
		UserId:   userId,
		Email:    email,
		UserRole: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "https://portfolioinstruments.com/api",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func VerifyJwt(tokenString string, jwtSecret string) (*JwtClaims, error) {
	t, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
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
