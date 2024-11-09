package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/alexedwards/argon2id"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Queries *sqlc.Queries
}

func (h *UserHandler) Create(c echo.Context) error {
	var user sqlc.CreateUserParams

	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	params := &argon2id.Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}

	saltedHash, err := argon2id.CreateHash(user.SaltedHash, params)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to hash password: %v", err))

		return echo.ErrInternalServerError
	}
	user.SaltedHash = saltedHash

	_, err = h.Queries.CreateUser(c.Request().Context(), user)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create user: %v", err))

		return echo.ErrInternalServerError
	}

	return c.String(http.StatusCreated, "User registered successfully")
}
