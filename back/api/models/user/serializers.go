package user

import (
	"context"
	"log/slog"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/internal"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserParams struct {
	Username  string `json:"username" validate:"required,min=5,max=20,alphanum"`
	Email     string `json:"email" validate:"required,email"`
	Firstname string `json:"firstname" validate:"required,max=20,alpha"`
	Lastname  string `json:"lastname" validate:"required,max=20,alpha"`
	Password  string `json:"password" validate:"required,min=5,max=20,ascii"`
}

type UpdateUserParams struct {
	Email     pgtype.Text `json:"email,omitempty" validate:"omitempty,email"`
	Firstname pgtype.Text `json:"firstname,omitempty" validate:"omitempty,max=20,alpha"`
	Lastname  pgtype.Text `json:"lastname,omitempty" validate:"omitempty,max=20,alpha"`
	Password  pgtype.Text `json:"password,omitempty" validate:"omitempty,min=5,max=20,ascii"`
}

func CreateUser(queries *sqlc.Queries, user CreateUserParams, isAdmin bool) error {
	c := context.Background()

	saltedHash, err := internal.CreateSaltedHash(user.Password)
	if err != nil {
		slog.Error("Failed to hash password",
			"error", err)
		return err
	}

	_, err = queries.CreateUser(c, sqlc.CreateUserParams{
		Username:   user.Username,
		Email:      user.Email,
		SaltedHash: saltedHash,
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		IsAdmin:    isAdmin,
	})
	if err != nil {
		slog.Error("Failed to create user",
			"error", err)
		return err
	}

	return nil
}
