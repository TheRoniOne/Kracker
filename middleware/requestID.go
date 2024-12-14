package middleware

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
)

func RequestIDMiddleware() echo.MiddlewareFunc {
	return middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			id, err := uuid.NewV7()
			if err != nil {
				slog.Error("Failed to generate UUID",
					"error", err)

				return random.String(32)
			}

			return id.String()
		},
	})
}
