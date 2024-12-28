package middleware

import (
	"log/slog"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type SessionMiddleware struct {
	Queries *sqlc.Queries
}

func (s *SessionMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("SESSION")
		if err != nil {
			return echo.ErrUnauthorized
		}

		sessionID := pgtype.UUID{}
		err = sessionID.Scan(cookie.Value)
		if err != nil {
			slog.Error("Failed to scan session ID",
				"error", err)
			return echo.ErrUnauthorized
		}

		session, err := s.Queries.GetSession(c.Request().Context(), sessionID)
		if err != nil {
			slog.Error("Failed to get session from database",
				"error", err)
			return echo.ErrUnauthorized
		}

		c.Set("UserID", session.UserID)

		return next(c)
	}
}
