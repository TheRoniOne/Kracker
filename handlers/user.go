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
		return c.String(http.StatusBadRequest, "bad request")
	}

	h.queries.CreateUser(c.Request().Context(), user)

	return c.String(http.StatusCreated, "User registered successfully")
}
