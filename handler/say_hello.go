package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SayHello(c echo.Context) error {
	name := c.QueryParam("name")

	if name == "" {
		name = "World"
	}

	return c.String(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}
