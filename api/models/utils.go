package models

import (
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func GetUserID(_ echo.Context) int64 {
	return 1 // TODO implement
}

func GetUserIDFromUser(u *sqlc.User) func(echo.Context) int64 {
	return func(_ echo.Context) int64 {
		return u.ID
	}
}

func GetSessionID(c echo.Context) pgtype.UUID {
	return c.Get("sessionID").(pgtype.UUID)
}
