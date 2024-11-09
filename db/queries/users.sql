-- name: GetUser :one
SELECT * FROM Users
WHERE id = $1 LIMIT 1;

-- name: GetUserFromUsername :one
SELECT * FROM Users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM Users;

-- name: CreateUser :one
INSERT INTO Users (
  username, email, salted_hash, firstname, lastname
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE Users
  set username = $2,
  email = $3,
  salted_hash = $4,
  firstname = $5,
  lastname = $6
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = $1;
