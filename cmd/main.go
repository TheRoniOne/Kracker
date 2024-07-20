package main

import (
	"github.com/TheRoniOne/Kracker/handler"
	"github.com/TheRoniOne/Kracker/initializer"
	"github.com/labstack/echo/v4"
)

func init() {
	initializer.LoadEnvVars()
	initializer.ConnectToDB()
}

func main() {
	e := echo.New()

	api := e.Group("/api")
	api.GET("/say-hello", handler.SayHello)

	e.Logger.Fatal(e.Start(":1323"))
}
