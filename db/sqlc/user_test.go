package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Email:             "user@example.org",
		EncryptedPassword: "qwertyuiop",
	}

	u, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.NotZero(t, u.ID)

	require.Equal(t, arg.Email, u.Email)
	require.Equal(t, arg.EncryptedPassword, u.EncryptedPassword)

	err = testQueries.DeleteUser(context.Background(), u.Email)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	arg := CreateUserParams{
		Email:             "user@example.org",
		EncryptedPassword: "qwertyuiop",
	}
	u, err := testQueries.CreateUser(context.Background(), arg)
	err = testQueries.DeleteUser(context.Background(), u.Email)

	u2, err := testQueries.GetUser(context.Background(), arg.Email)
	require.Error(t, err)
	require.Empty(t, u2)
}

func TestGetUser(t *testing.T) {
	arg := CreateUserParams{
		Email:             "user@example.org",
		EncryptedPassword: "qwertyuiop",
	}
	u1, _ := testQueries.CreateUser(context.Background(), arg)

	u2, err := testQueries.GetUser(context.Background(), arg.Email)
	require.NoError(t, err)

	require.Equal(t, u1.Email, u2.Email)
	require.Equal(t, u1.EncryptedPassword, u2.EncryptedPassword)
	require.Equal(t, u1.ID, u2.ID)

	require.NotEmpty(t, u2)
	require.NotZero(t, u2.ID)

	err = testQueries.DeleteUser(context.Background(), u1.Email)
	require.NoError(t, err)
}

func TestGetUserForUpdate(t *testing.T) {
	arg := CreateUserParams{
		Email:             "user@example.org",
		EncryptedPassword: "qwertyuiop",
	}
	u1, _ := testQueries.CreateUser(context.Background(), arg)

	u2, err := testQueries.GetUserForUpdate(context.Background(), arg.Email)
	require.NoError(t, err)

	require.Equal(t, u1.Email, u2.Email)
	require.Equal(t, u1.EncryptedPassword, u2.EncryptedPassword)
	require.Equal(t, u1.ID, u2.ID)

	require.NotEmpty(t, u2)
	require.NotZero(t, u2.ID)

	err = testQueries.DeleteUser(context.Background(), u1.Email)
	require.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	arg1 := CreateUserParams{
		Email:             "user@example.org",
		EncryptedPassword: "qwertyuiop",
	}
	arg2 := UpdateUserParams{
		Email:             "user@example.org",
		EncryptedPassword: "qwertyuiop1",
	}
	u1, err := testQueries.CreateUser(context.Background(), arg1)

	u2, err := testQueries.UpdateUser(context.Background(), arg2)

	require.NoError(t, err)
	require.Equal(t, u1.Email, u2.Email)
	require.Equal(t, u2.EncryptedPassword, arg2.EncryptedPassword)

	require.NotEmpty(t, u2)
	require.NotZero(t, u2.ID)

	err = testQueries.DeleteUser(context.Background(), u2.Email)
	require.NoError(t, err)
}
