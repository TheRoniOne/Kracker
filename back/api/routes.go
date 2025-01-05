package api

import (
	"net/http"

	"github.com/TheRoniOne/Kracker/api/models/session"
	"github.com/TheRoniOne/Kracker/api/models/user"
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

	sessionHandler := session.Handler{Queries: queries}
	group.POST("/session", sessionHandler.Create)
	group.DELETE("/session", sessionHandler.Delete)

	userHandler := user.Handler{Queries: queries}
	group.POST("/user", userHandler.Create)
	protected.GET("/user/list", userHandler.List)
	protected.PATCH("/user", userHandler.Update)
}
