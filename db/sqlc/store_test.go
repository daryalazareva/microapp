package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTempUser(t *testing.T) User {
	arg := CreateUserParams{
		Email:             "user@example.org",
		EncryptedPassword: "qwertyuiop",
	}
	u, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	return u
}

func deleteTempUser(t *testing.T) {
	email := "user@example.org"
	err := testQueries.DeleteUser(context.Background(), email)
	if err != nil {
		log.Fatal(err)
	}
}

func TestChangePasswordTx(t *testing.T) {
	store := NewStore(testDB)
	u := createTempUser(t)
	newpassword := "qwertyuiopa"

	err := store.ChangePasswordTx(context.Background(), u.Email, newpassword)
	require.NoError(t, err)

	//run concurrent cases
	n := 49
	errs1 := make(chan error)
	errs2 := make(chan error)

	for i := 0; i < n; i++ {
		go func() {
			err := store.ChangePasswordTx(context.Background(), u.Email, newpassword)
			errs1 <- err
		}()
		go func() {
			_, err := store.GetRecordUser(context.Background(), u.Email)
			errs2 <- err
		}()
	}

	//check result
	for i := 0; i < n; i++ {
		err := <-errs1
		require.NoError(t, err)
		err = <-errs2
		require.NoError(t, err)
	}

	deleteTempUser(t)
}

func TestCreateRecordUser(t *testing.T) {
	store := NewStore(testDB)
	arg := CreateUserParams{
		Email:             "user@example.org",
		EncryptedPassword: "qwertyuiop",
	}

	u1, err := store.CreateRecordUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, u1)
	require.NotZero(t, u1.ID)

	u2, err := store.GetRecordUser(context.Background(), arg.Email)
	require.Equal(t, u1.ID, u2.ID)
	require.Equal(t, u1.Email, u2.Email)
	require.Equal(t, u1.EncryptedPassword, u2.EncryptedPassword)

	deleteTempUser(t)
}

func TestGetRecordUser(t *testing.T) {
	store := NewStore(testDB)
	u1 := createTempUser(t)

	u2, err := store.GetRecordUser(context.Background(), u1.Email)
	require.NoError(t, err)
	require.Equal(t, u2.Email, u1.Email)
	require.Equal(t, u2.EncryptedPassword, u1.EncryptedPassword)
	require.Equal(t, u2.ID, u1.ID)
	require.NotEmpty(t, u2)
	require.NotZero(t, u2.ID)

	deleteTempUser(t)
}
