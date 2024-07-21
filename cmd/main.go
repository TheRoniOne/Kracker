package main

import (
	"os"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/handlers"
	"github.com/TheRoniOne/Kracker/initializers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

var dBPool *pgxpool.Pool

func init() {
	initializers.LoadEnvVars()
	dBPool = initializers.ConnectToDB(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
}

func main() {
	defer dBPool.Close()

	e := echo.New()
	queries := sqlc.New(dBPool)
	handlers.SetUpRoutes(e, queries)

	e.Logger.Fatal(e.Start(":1323"))
}
