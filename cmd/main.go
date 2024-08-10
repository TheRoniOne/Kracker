package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/handlers"
	"github.com/TheRoniOne/Kracker/internal"
	"github.com/TheRoniOne/Kracker/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

var (
	dbPool *pgxpool.Pool
	logger *slog.Logger
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	logger = slog.Default()

	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", internal.DB_USER, internal.DB_PASSWORD, internal.DB_HOST, internal.DB_PORT, internal.DB_NAME)
	dbPool, err = pgxpool.New(context.Background(), connStr)

	if err != nil {
		logger.Error(fmt.Sprintf("Unable to create connection pool: %v", err))
		panic(err)
	}
}

func main() {
	defer dbPool.Close()

	e := echo.New()

	e.Debug = internal.DEBUG
	e.Use(middleware.LoggingMiddleware(logger))

	queries := sqlc.New(dbPool)
	handlers.SetUpRoutes(e, queries)

	e.Logger.Fatal(e.Start(":1323"))
}
