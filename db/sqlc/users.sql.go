// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO Users (
  username, email, salted_hash, firstname, lastname
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, username, email, salted_hash, firstname, lastname
`

type CreateUserParams struct {
	Username   string
	Email      string
	SaltedHash string
	Firstname  string
	Lastname   string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.SaltedHash,
		arg.Firstname,
		arg.Lastname,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.SaltedHash,
		&i.Firstname,
		&i.Lastname,
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
SELECT id, username, email, salted_hash, firstname, lastname FROM Users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.SaltedHash,
		&i.Firstname,
		&i.Lastname,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, salted_hash, firstname, lastname FROM Users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.SaltedHash,
			&i.Firstname,
			&i.Lastname,
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

const updateUser = `-- name: UpdateUser :exec
UPDATE Users
  set username = $2,
  email = $3,
  salted_hash = $4,
  firstname = $5,
  lastname = $6
WHERE id = $1
`

type UpdateUserParams struct {
	ID         int64
	Username   string
	Email      string
	SaltedHash string
	Firstname  string
	Lastname   string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.SaltedHash,
		arg.Firstname,
		arg.Lastname,
	)
	return err
}
