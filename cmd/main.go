package main

import (
	"github.com/TheRoniOne/Kracker/handler"
	"github.com/TheRoniOne/Kracker/initializers"
	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	e := echo.New()

	api := e.Group("/api")
	api.GET("/say-hello", handler.SayHello)

	e.Logger.Fatal(e.Start(":1323"))
}
