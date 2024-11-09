package api

import (
	"github.com/TheRoniOne/Kracker/api/models"
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/labstack/echo/v4"
)

func SetUpRoutes(app *echo.Echo, queries *sqlc.Queries) {
	group := app.Group("/api")

	userHandler := models.UserHandler{Queries: queries, GetUserID: models.GetUserID}
	group.POST("/user/create", userHandler.Create)
	group.GET("/user/list", userHandler.List)
}
