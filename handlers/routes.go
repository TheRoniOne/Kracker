package handlers

import (
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/labstack/echo/v4"
)

func SetUpRoutes(app *echo.Echo, queries *sqlc.Queries) {
	group := app.Group("/api")
	group.GET("/say-hello", SayHello)

	userHandler := UserHandler{Queries: queries}
	group.POST("/user/create", userHandler.Create)
}
