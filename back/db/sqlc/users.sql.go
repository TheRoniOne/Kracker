// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO Users (
  username, email, salted_hash, firstname, lastname, is_admin
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, username, email, salted_hash, firstname, lastname, is_admin
`

type CreateUserParams struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	SaltedHash string `json:"salted_hash"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	IsAdmin    bool   `json:"is_admin"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.SaltedHash,
		arg.Firstname,
		arg.Lastname,
		arg.IsAdmin,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.SaltedHash,
		&i.Firstname,
		&i.Lastname,
		&i.IsAdmin,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, firstname, lastname, is_admin FROM Users
WHERE id = $1 LIMIT 1
`

type GetUserRow struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	IsAdmin   bool   `json:"is_admin"`
}

func (q *Queries) GetUser(ctx context.Context, id int64) (GetUserRow, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
		&i.IsAdmin,
	)
	return i, err
}

const getUserFromUsername = `-- name: GetUserFromUsername :one
SELECT id, username, email, salted_hash, firstname, lastname, is_admin FROM Users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserFromUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserFromUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.SaltedHash,
		&i.Firstname,
		&i.Lastname,
		&i.IsAdmin,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, firstname, lastname, is_admin FROM Users
`

type ListUsersRow struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	IsAdmin   bool   `json:"is_admin"`
}

func (q *Queries) ListUsers(ctx context.Context) ([]ListUsersRow, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersRow
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Firstname,
			&i.Lastname,
			&i.IsAdmin,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE Users
  set updated_at = now(),
  username = COALESCE($1, username),
  email = COALESCE($2, email),
  salted_hash = COALESCE($3, salted_hash),
  firstname = COALESCE($4, firstname),
  lastname = COALESCE($5, lastname),
  is_admin = COALESCE($6, is_admin)
WHERE id = $7
RETURNING id, username, email, salted_hash, firstname, lastname, is_admin
`

type UpdateUserParams struct {
	Username   pgtype.Text `json:"username"`
	Email      pgtype.Text `json:"email"`
	SaltedHash pgtype.Text `json:"salted_hash"`
	Firstname  pgtype.Text `json:"firstname"`
	Lastname   pgtype.Text `json:"lastname"`
	IsAdmin    pgtype.Bool `json:"is_admin"`
	ID         int64       `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Username,
		arg.Email,
		arg.SaltedHash,
		arg.Firstname,
		arg.Lastname,
		arg.IsAdmin,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.SaltedHash,
		&i.Firstname,
		&i.Lastname,
		&i.IsAdmin,
	)
	return i, err
}
