package utils

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/TheRoniOne/Kracker/internal"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var Migrations []string

const (
	testDBName     = "test-db"
	testDBUser     = "postgres"
	testDBPassword = "postgres123"
)

func init() {
	var err error
	migrationsPath := filepath.Join(internal.RootPath, "db/migrations/*.sql")
	Migrations, err = filepath.Glob(migrationsPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while trying to find migrations: %v\n", err)
		os.Exit(1)
	}
}

func SetUpTestDBPool(ctx context.Context, t *testing.T) *pgxpool.Pool {
	pgContainer, err := postgres.Run(ctx,
		"postgres:17-alpine",
		postgres.WithInitScripts(Migrations...),
		postgres.WithDatabase(testDBName),
		postgres.WithUsername(testDBUser),
		postgres.WithPassword(testDBPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("Failed to terminate pgContainer: %s", err)
		}
	})

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	dbPool, err := pgxpool.New(context.Background(), connStr)
	require.NoError(t, err)

	return dbPool
}
