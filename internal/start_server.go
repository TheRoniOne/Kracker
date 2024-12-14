package internal

import (
	"fmt"
	"log/slog"
	"net"
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

func StartTestServer(e *echo.Echo) string {
	port, err := getUnusedPort()
	if err != nil {
		slog.Error("Failed to get unused port",
			"error", err)
		return ""
	}

	StartServer(e, fmt.Sprintf(":%d", port))

	time.Sleep(1 * time.Second)

	return fmt.Sprintf("http://localhost:%d", port)
}

func getUnusedPort() (int, error) {
	listener, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port, nil
}
