package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"

	"github.com/TheRoniOne/Kracker/api"
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/internal"
	"github.com/TheRoniOne/Kracker/middleware"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

var (
	dbPool *pgxpool.Pool
	logger *slog.Logger
)

func init() {
	var slogOpts *slog.HandlerOptions
	if internal.Debug {
		slogOpts = &slog.HandlerOptions{Level: slog.LevelDebug}
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, slogOpts)))

	logger = slog.Default()

	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", internal.DBUser, url.QueryEscape(internal.DBPassword), internal.DBHost, internal.DBPort, internal.DBName)
	dbPool, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		logger.Error(fmt.Sprintf("Unable to create connection pool: %v", err))
		panic(err)
	}
}

func main() {
	defer dbPool.Close()

	e := echo.New()

	e.Debug = internal.Debug
	e.Use(echomiddleware.RateLimiter(echomiddleware.NewRateLimiterMemoryStore(rate.Limit(internal.RateLimit))))
	e.Use(middleware.RequestIDMiddleware(logger))
	e.Use(middleware.LoggingMiddleware(logger))
	e.Use(echomiddleware.Recover())

	queries := sqlc.New(dbPool)
	api.SetUpRoutes(e, queries)

	e.Logger.Fatal(e.Start(":1323"))
}
