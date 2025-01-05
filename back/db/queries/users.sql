-- name: GetUser :one
SELECT id, username, email, firstname, lastname, is_admin FROM Users
WHERE id = $1 LIMIT 1;

-- name: GetUserFromUsername :one
SELECT * FROM Users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT id, username, email, firstname, lastname, is_admin FROM Users;

-- name: CreateUser :one
INSERT INTO Users (
  username, email, salted_hash, firstname, lastname, is_admin
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateUser :one
UPDATE Users
  set updated_at = now(),
  username = COALESCE(sqlc.narg('username'), username),
  email = COALESCE(sqlc.narg('email'), email),
  salted_hash = COALESCE(sqlc.narg('salted_hash'), salted_hash),
  firstname = COALESCE(sqlc.narg('firstname'), firstname),
  lastname = COALESCE(sqlc.narg('lastname'), lastname),
  is_admin = COALESCE(sqlc.narg('is_admin'), is_admin)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = $1;
