package main

import (
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/handlers"
	"github.com/TheRoniOne/Kracker/initializers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

var DBPool *pgxpool.Pool

func init() {
	initializers.LoadEnvVars()
	DBPool := initializers.ConnectToDB()
}

func main() {
	defer DBPool.Close()
	e := echo.New()

	queries := sqlc.New(DBPool)

	handlers.SetUpRoutes(e, queries)

	e.Logger.Fatal(e.Start(":1323"))
}
