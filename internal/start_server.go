package internal

import (
	"fmt"
	"net"
	"time"

	"github.com/TheRoniOne/Kracker/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func StartServer(e *echo.Echo, address string) error {
	e.Debug = Debug
	e.Use(middleware.RequestIDMiddleware())
	e.Use(middleware.LoggingMiddleware())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.RateLimiter(echomiddleware.NewRateLimiterMemoryStore(rate.Limit(RateLimit))))
	e.Use(echomiddleware.TimeoutWithConfig(echomiddleware.TimeoutConfig{
		Timeout: 25 * time.Second,
	}))

	return e.Start(address)
}

func StartTestServer(e *echo.Echo) error {
	port, err := getUnusedPort()
	if err != nil {
		return err
	}

	return StartServer(e, fmt.Sprintf(":%d", port))
}

func getUnusedPort() (int, error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port, nil
}
