package internal

import (
	"context"
	"log/slog"

	"github.com/TheRoniOne/Kracker/db/sqlc"
)

func CreateUser(user sqlc.CreateUserParams, queries *sqlc.Queries) error {
	c := context.Background()

	saltedHash, err := CreateSaltedHash(user.SaltedHash)
	if err != nil {
		slog.Error("Failed to hash password",
			"error", err)
		return err
	}
	user.SaltedHash = saltedHash

	_, err = queries.CreateUser(c, user)
	if err != nil {
		slog.Error("Failed to create user",
			"error", err)
		return err
	}

	return nil
}
