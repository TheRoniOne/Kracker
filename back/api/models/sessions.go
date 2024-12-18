package models

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/internal"
	"github.com/alexedwards/argon2id"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type SessionHandler struct {
	Queries *sqlc.Queries
}

type SessionCreateParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *SessionHandler) Create(c echo.Context) error {
	var loginParams SessionCreateParams
	err := c.Bind(&loginParams)
	if err != nil {
		return echo.ErrBadRequest
	}

	userData, err := h.Queries.GetUserFromUsername(c.Request().Context(), loginParams.Username)
	if err != nil {
		return err
	}

	isOk, err := argon2id.ComparePasswordAndHash(loginParams.Password, userData.SaltedHash)
	if err != nil {
		return err
	}

	if !isOk {
		return echo.ErrUnauthorized
	}

	tStamp := pgtype.Timestamptz{}

	now := time.Now().AddDate(0, 0, internal.SessionMaxAgeDays).In(internal.TimeLocation)
	err = tStamp.Scan(now)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to scan time: %v", err))
	}

	session, err := h.Queries.CreateSession(c.Request().Context(), sqlc.CreateSessionParams{
		ExpiresAt: pgtype.Timestamptz{},
		UserID:    userData.ID,
	})
	if err != nil {
		return err
	}

	sessID, err := session.ID.MarshalJSON()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get session ID: %v", err))
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, sessID)
}

func (h *SessionHandler) Delete(c echo.Context) error {
	err := h.Queries.DeleteSession(c.Request().Context(), GetSessionID(c))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to delete session: %v", err))
		return echo.ErrInternalServerError
	}

	return c.String(http.StatusOK, "Logged out successfully")
}
