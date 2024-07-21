package main

import (
	"github.com/TheRoniOne/Kracker/handlers"
	"github.com/TheRoniOne/Kracker/initializers"
	"github.com/labstack/echo/v4"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	defer initializers.DBPool.Close()
	e := echo.New()

	api := e.Group("/api")
	api.GET("/say-hello", handlers.SayHello)

	e.Logger.Fatal(e.Start(":1323"))
}
