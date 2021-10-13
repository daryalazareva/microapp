// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email, encrypted_password
) VALUES (
  $1, $2
)
RETURNING id, email, encrypted_password
`

type CreateUserParams struct {
	Email             string `json:"email"`
	EncryptedPassword string `json:"encrypted_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.EncryptedPassword)
	var i User
	err := row.Scan(&i.ID, &i.Email, &i.EncryptedPassword)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1
`

func (q *Queries) DeleteUser(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, email)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, email, encrypted_password FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(&i.ID, &i.Email, &i.EncryptedPassword)
	return i, err
}

const getUserForUpdate = `-- name: GetUserForUpdate :one
SELECT id, email, encrypted_password FROM users
WHERE email = $1 FOR UPDATE
`

func (q *Queries) GetUserForUpdate(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserForUpdate, email)
	var i User
	err := row.Scan(&i.ID, &i.Email, &i.EncryptedPassword)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET encrypted_password = $2
WHERE email = $1
RETURNING id, email, encrypted_password
`

type UpdateUserParams struct {
	Email             string `json:"email"`
	EncryptedPassword string `json:"encrypted_password"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Email, arg.EncryptedPassword)
	var i User
	err := row.Scan(&i.ID, &i.Email, &i.EncryptedPassword)
	return i, err
}