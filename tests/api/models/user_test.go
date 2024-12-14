package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/TheRoniOne/Kracker/api"
	"github.com/TheRoniOne/Kracker/api/models"
	"github.com/TheRoniOne/Kracker/db/factories"
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/tests/utils"
	"github.com/alexedwards/argon2id"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserCreate(t *testing.T) {
	ctx := context.Background()
	connStr := utils.SetUpTestWithDB(ctx, t)

	dbPool, err := pgxpool.New(context.Background(), connStr)
	require.NoError(t, err)

	e := echo.New()
	queries := sqlc.New(dbPool)

	serverURL := utils.StartTestServer(e)
	require.NotEmpty(t, serverURL)
	defer e.Close()

	api.SetUpRoutes(e, queries)

	userData := models.UserCreateParams{
		Username:  "test",
		Email:     "test",
		Password:  "test",
		Firstname: "test",
		Lastname:  "test",
	}

	j, _ := json.Marshal(userData)

	response, err := http.Post(serverURL+"/api/user/create", echo.MIMEApplicationJSON, strings.NewReader(string(j)))
	require.NoError(t, err)

	if assert.Equal(t, http.StatusCreated, response.StatusCode) {
		users, _ := queries.ListUsers(ctx)
		assert.Len(t, users, 1)

		user := users[0]
		assert.Equal(t, userData.Username, user.Username)
		assert.Equal(t, userData.Email, user.Email)
		assert.Equal(t, userData.Firstname, user.Firstname)
		assert.Equal(t, userData.Lastname, user.Lastname)
		assert.False(t, user.IsAdmin)

		userDBData, _ := queries.GetUserFromUsername(ctx, userData.Username)
		isOk, _ := argon2id.ComparePasswordAndHash(userData.Password, userDBData.SaltedHash)
		assert.True(t, isOk)
	}
}

func TestUserList(t *testing.T) {
	ctx := context.Background()
	connStr := utils.SetUpTestWithDB(ctx, t)

	dbPool, err := pgxpool.New(context.Background(), connStr)
	require.NoError(t, err)

	e := echo.New()
	queries := sqlc.New(dbPool)

	serverURL := utils.StartTestServer(e)
	require.NotEmpty(t, serverURL)
	defer e.Close()

	api.SetUpRoutes(e, queries)

	UserFactory := factories.UserFactory{Queries: queries}
	UserFactory.CreateOne()

	response, err := http.Get(serverURL + "/api/user/list")
	require.NoError(t, err)

	if assert.Equal(t, http.StatusOK, response.StatusCode) {

		users, _ := queries.ListUsers(ctx)

		expected, _ := json.Marshal(users)

		assert.Equal(t, string(expected)+"\n", utils.ReadRespBody(response))
	}
}
