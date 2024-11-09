package models

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/TheRoniOne/Kracker/db/factories"
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/internal"
	"github.com/alexedwards/argon2id"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserCreate(t *testing.T) {
	ctx := context.Background()
	connStr := internal.SetUpTestWithDB(ctx, t)

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

	UserFactory := factories.UserFactory{Queries: queries}
	user := UserFactory.CreateOne()

	handler := UserHandler{Queries: queries, GetUserID: GetUserIDFromUser(&user)}
	if assert.NoError(t, handler.Create(echoContext)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		users, _ := queries.ListUsers(ctx)
		assert.Len(t, users, 2)

		user := users[1]
		assert.Equal(t, userData.Username, user.Username)
		assert.Equal(t, userData.Email, user.Email)
		assert.Equal(t, userData.Firstname, user.Firstname)
		assert.Equal(t, userData.Lastname, user.Lastname)

		userDBData, _ := queries.GetUserFromUsername(ctx, userData.Username)
		isOk, _ := argon2id.ComparePasswordAndHash(userData.SaltedHash, userDBData.SaltedHash)
		assert.True(t, isOk)
	}
}

func TestUserList(t *testing.T) {
	ctx := context.Background()
	connStr := internal.SetUpTestWithDB(ctx, t)

	dbPool, err := pgxpool.New(context.Background(), connStr)
	require.NoError(t, err)

	e := echo.New()
	queries := sqlc.New(dbPool)

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string("")))
	rec := httptest.NewRecorder()
	echoContext := e.NewContext(req, rec)

	UserFactory := factories.UserFactory{Queries: queries}
	user := UserFactory.CreateOne()

	handler := UserHandler{Queries: queries, GetUserID: GetUserIDFromUser(&user)}
	if assert.NoError(t, handler.List(echoContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		users, _ := queries.ListUsers(ctx)

		expected, _ := json.Marshal(users)
		assert.Equal(t, string(expected)+"\n", rec.Body.String())
	}
}
