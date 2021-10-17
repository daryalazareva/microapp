package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := "123456"

	encryptedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, encryptedPassword)

	err = CheckPassword(password, encryptedPassword)
	require.NoError(t, err)

	wrongPassword := "456789"
	err = CheckPassword(wrongPassword, encryptedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
