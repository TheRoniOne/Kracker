package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/TheRoniOne/Kracker/db"
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestUserCreate(t *testing.T) {
	ctx := context.Background()
	dbName := "test-db"
	dbUser := "postgres"
	dbPassword := "postgres123"

	pgContainer, err := postgres.Run(ctx,
		"postgres:17-alpine",
		postgres.WithInitScripts(db.Migrations...),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
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

	e := echo.New()
	queries := sqlc.New(dbPool)

	userData := sqlc.CreateUserParams{
		Username:   "test",
		Email:      "test",
		SaltedHash: "test",
		Firstname:  "test",
		Lastname:   "test",
	}

	j, _ := json.Marshal(userData)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(j)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	echoContext := e.NewContext(req, rec)

	handler := UserHandler{Queries: queries}
	if assert.NoError(t, handler.Create(echoContext)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		users, _ := queries.ListUsers(ctx)
		assert.Len(t, users, 1)

		user := users[0]
		assert.Equal(t, userData.Username, user.Username)
		assert.Equal(t, userData.Email, user.Email)
		assert.Equal(t, userData.SaltedHash, user.SaltedHash)
		assert.Equal(t, userData.Firstname, user.Firstname)
		assert.Equal(t, userData.Lastname, user.Lastname)
	}
}
