package helper

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}

	return string(hashedPassword), nil
}

func CheckPassword(hashPass, dbPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(dbPass))

	return err == nil
}