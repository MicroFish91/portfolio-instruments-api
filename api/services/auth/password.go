package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	pbytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pbytes), nil
}

func CompareHashAndPassword(encpw string, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw))
}
