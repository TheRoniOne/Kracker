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

func (h *UserHandler) Create(c echo.Context) error {
	var user internal.CreateUserParams

	err := c.Bind(&user)
	if err != nil {
		return echo.ErrBadRequest
	}

	err = c.Validate(user)
	if err != nil {
		slog.Error("Failed to validate create user params",
			"error", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	err = internal.CreateUser(h.Queries, user, false)
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
