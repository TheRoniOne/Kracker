package internal

import (
	"context"
	"log/slog"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/alexedwards/argon2id"
)

func CreateUser(user sqlc.CreateUserParams, queries *sqlc.Queries) error {
	c := context.Background()

	params := &argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	saltedHash, err := argon2id.CreateHash(user.SaltedHash, params)
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
