package models

import (
	"github.com/labstack/echo/v4"
)

func GetUserID(c echo.Context) int64 {
	return c.Get("UserID").(int64)
}
