package api

import (
	"net/http"

	"github.com/TheRoniOne/Kracker/api/models"
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/middleware"
	"github.com/labstack/echo/v4"
)

func SetUpRoutes(app *echo.Echo, queries *sqlc.Queries) {
	app.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	group := app.Group("/api")

	sessionMiddleware := middleware.SessionMiddleware{Queries: queries}
	protected := group.Group("", sessionMiddleware.Handle)

	sessionHandler := models.SessionHandler{Queries: queries}
	group.POST("/session", sessionHandler.Create)
	group.DELETE("/session", sessionHandler.Delete)

	userHandler := models.UserHandler{Queries: queries}
	group.POST("/user/create", userHandler.Create)
	protected.GET("/user/list", userHandler.List)
}
