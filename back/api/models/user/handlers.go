package user

import (
	"log/slog"
	"net/http"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/internal"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Queries *sqlc.Queries
}

func (h *Handler) Create(c echo.Context) error {
	var user CreateUserParams

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

	err = CreateUser(h.Queries, user, false)
	if err != nil {
		return err
	}

	return c.String(http.StatusCreated, "User registered successfully")
}

func (h *Handler) List(c echo.Context) error {
	users, err := h.Queries.ListUsers(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (h *Handler) Update(c echo.Context) error {
	var params UpdateUserParams

	err := c.Bind(&params)
	if err != nil {
		slog.Error("Failed to bind params",
			"error", err)
		return echo.ErrBadRequest
	}

	err = c.Validate(params)
	if err != nil {
		slog.Error("Failed to validate user update params",
			"error", err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	var saltedHash pgtype.Text

	if params.Password.Valid {
		password := params.Password.String

		s, err := internal.CreateSaltedHash(password)
		if err != nil {
			slog.Error("Failed to hash password",
				"error", err)
			return echo.ErrInternalServerError
		}

		saltedHash = pgtype.Text{}
		err = saltedHash.Scan(s)
		if err != nil {
			slog.Error("Failed to scan salted hash",
				"error", err)
			return echo.ErrInternalServerError
		}
	}

	updatedUser, err := h.Queries.UpdateUser(c.Request().Context(), sqlc.UpdateUserParams{
		Email:      params.Email,
		Firstname:  params.Firstname,
		Lastname:   params.Lastname,
		SaltedHash: saltedHash,
	})
	if err != nil {
		slog.Error("Failed to update user",
			"error", err)
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, updatedUser)
}
