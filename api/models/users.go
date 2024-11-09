package models

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/alexedwards/argon2id"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Queries   *sqlc.Queries
	GetUserID func(c echo.Context) int64
}

func (h *UserHandler) Create(c echo.Context) error {
	var user sqlc.CreateUserParams

	err := c.Bind(&user)
	if err != nil {
		return echo.ErrBadRequest
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
		return err
	}
	user.SaltedHash = saltedHash

	_, err = h.Queries.CreateUser(c.Request().Context(), user)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "User registered successfully")
}

func (h *UserHandler) List(c echo.Context) error {
	users, err := h.Queries.ListUsers(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
