package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(encryptedPassword), nil
}

//checks if the provided password is correct
func CheckPassword(password string, encryptedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
}
