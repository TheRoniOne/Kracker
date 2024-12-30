package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"time"

	"github.com/TheRoniOne/Kracker/api"
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/internal"

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

	exitChannel := make(chan bool)

	e := echo.New()
	internal.StartServer(e, ":1323", exitChannel)

	queries := sqlc.New(dbPool)
	api.SetUpRoutes(e, queries)

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()

		for range ticker.C {
			queries.DeleteExpiredSessions(context.Background())
		}
	}()

	<-exitChannel
}
