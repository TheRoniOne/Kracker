package handlers

import (
	"net/http"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	queries *sqlc.Queries
}

func (h *UserHandler) Register(c echo.Context) error {
	var user sqlc.CreateUserParams

	err := c.Bind(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}

	_, err = h.queries.CreateUser(c.Request().Context(), user)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "User registered successfully")
}
