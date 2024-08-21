package utils

import "github.com/MicroFish91/portfolio-instruments-api/api/services/auth"

func Generate40xTokens(userId int, email string) (tok401 string, tok403 string, err error) {
	t401, err := auth.GenerateSignedJwt(userId, email, "Default")
	if err != nil {
		return "", "", err
	}
	t401 = t401[1:]

	t403, err := auth.GenerateSignedJwt(100, "fake_user_100@gmail.com", "Default")
	if err != nil {
		return "", "", err
	}

	return t401, t403, nil
}
