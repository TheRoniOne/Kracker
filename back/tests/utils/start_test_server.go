package utils

import (
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/TheRoniOne/Kracker/internal"
	"github.com/labstack/echo/v4"
)

func StartTestServer(e *echo.Echo) string {
	port, err := getUnusedPort()
	if err != nil {
		slog.Error("Failed to get unused port",
			"error", err)
		return ""
	}

	internal.StartServer(e, fmt.Sprintf(":%d", port), make(chan bool))

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
