package main

import (
	"github.com/TheRoniOne/Kracker/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	api := e.Group("/api")
	api.GET("/say-hello", handler.SayHello)

	e.Logger.Fatal(e.Start(":1323"))
}
