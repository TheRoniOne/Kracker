package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"
	"os"

	"github.com/TheRoniOne/Kracker/api/models"
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/internal"
	"github.com/jackc/pgx/v5/pgxpool"
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
	fmt.Println("Enter the details for the superuser account:")
	username := internal.GetInput("Username")
	email := internal.GetInput("Email")
	password := internal.GetInput("Password")

	user := models.CreateUserParams{
		Username: username,
		Email:    email,
		Password: password,
	}

	queries := sqlc.New(dbPool)
	err := models.CreateUser(queries, user, true)
	if err != nil {
		slog.Error("Failed to create superuser",
			"error", err)
	}
}
