package internal

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/TheRoniOne/Kracker/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func StartServer(e *echo.Echo, address string, exit chan bool) {
	e.Debug = Debug

	e.Use(echomiddleware.CSRFWithConfig(echomiddleware.CSRFConfig{
		TokenLookup:    "cookie:_csrf",
		CookiePath:     "/",
		CookieDomain:   DOMAIN,
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
	}))
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.RateLimiter(echomiddleware.NewRateLimiterMemoryStore(rate.Limit(RateLimit))))
	e.Use(echomiddleware.TimeoutWithConfig(echomiddleware.TimeoutConfig{
		Timeout: 25 * time.Second,
	}))
	e.Use(middleware.RequestIDMiddleware())
	e.Use(middleware.LoggingMiddleware())

	go func() {
		err := e.Start(address)
		if err != nil {
			slog.Info("Server is down",
				"error", err)
		}

		exit <- true
	}()
}
