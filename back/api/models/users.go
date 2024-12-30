package models

import (
	"log/slog"
	"net/http"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/internal"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Queries *sqlc.Queries
}

type UserCreateParams struct {
	Username  string `json:"username" validate:"required,min=5,max=20,alphanum"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=5,max=20,alphanum"`
	Firstname string `json:"firstname" validate:"required,max=20,alpha"`
	Lastname  string `json:"lastname" validate:"required,max=20,alpha"`
}

func (h *UserHandler) Create(c echo.Context) error {
	var user UserCreateParams

	err := c.Bind(&user)
	if err != nil {
		return echo.ErrBadRequest
	}

	err = c.Validate(user)
	if err != nil {
		slog.Error("Failed to validate user create params",
			"error", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	saltedHash, err := internal.CreateSaltedHash(user.Password)
	if err != nil {
		slog.Error("Failed to hash password",
			"error", err)
		return echo.ErrInternalServerError
	}

	_, err = h.Queries.CreateUser(c.Request().Context(), sqlc.CreateUserParams{
		Username:   user.Username,
		Email:      user.Email,
		SaltedHash: saltedHash,
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
	})
	if err != nil {
		slog.Error("Failed to create user",
			"error", err)
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
