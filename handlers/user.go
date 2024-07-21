package handlers

import (
	"net/http"

	"github.com/TheRoniOne/Kracker/db/sqlc"
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

	_, err = h.Queries.CreateUser(c.Request().Context(), user)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "User registered successfully")
}
