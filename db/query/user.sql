-- name: CreateUser :one
INSERT INTO users (
  email, encrypted_password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users SET encrypted_password = $2
WHERE email = $1
RETURNING *;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE email = $1 FOR UPDATE;