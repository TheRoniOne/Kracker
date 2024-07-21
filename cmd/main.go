package main

import (
	"context"
	"fmt"
	"os"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/handlers"
	"github.com/TheRoniOne/Kracker/initializers"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

var dbPool *pgxpool.Pool

func init() {
	var err error
	initializers.LoadEnvVars()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	dbPool, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		panic(fmt.Errorf("unable to create connection pool: %v", err))
	}
}

func main() {
	defer dbPool.Close()

	e := echo.New()
	e.Debug = os.Getenv("DEBUG") == "true"

	queries := sqlc.New(dbPool)
	handlers.SetUpRoutes(e, queries)

	e.Logger.Fatal(e.Start(":1323"))
}
