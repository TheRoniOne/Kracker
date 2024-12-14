package internal

import (
	"log/slog"
	"time"

	"github.com/TheRoniOne/Kracker/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func StartServer(e *echo.Echo, address string) {
	e.Debug = Debug
	e.Use(middleware.RequestIDMiddleware())
	e.Use(middleware.LoggingMiddleware())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.RateLimiter(echomiddleware.NewRateLimiterMemoryStore(rate.Limit(RateLimit))))
	e.Use(echomiddleware.TimeoutWithConfig(echomiddleware.TimeoutConfig{
		Timeout: 25 * time.Second,
	}))

	go func() {
		err := e.Start(address)
		if err != nil {
			slog.Error("Server is down", "error", err)
		}
	}()
}
