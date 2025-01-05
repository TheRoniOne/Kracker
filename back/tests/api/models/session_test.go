package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/TheRoniOne/Kracker/api"
	"github.com/TheRoniOne/Kracker/api/models/session"
	"github.com/TheRoniOne/Kracker/db/builders"
	"github.com/TheRoniOne/Kracker/db/sqlc"
	"github.com/TheRoniOne/Kracker/tests/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSessionCreate(t *testing.T) {
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

	UserBuilder := builders.NewUserBuilder(queries).Username("test").Password("test")
	UserBuilder.CreateOne()

	createSessionParams := session.CreateParams{
		Username: "test",
		Password: "test",
	}

	body, err := json.Marshal(createSessionParams)
	require.NoError(t, err)

	apiClient := utils.NewLoggedInAPIClient(serverURL, "test", "test")
	response, err := apiClient.Post("/api/session", echo.MIMEApplicationJSON, body)
	require.NoError(t, err)

	if assert.Equal(t, http.StatusOK, response.StatusCode) {
		sessionCookie := response.Cookies()[1]

		assert.Equal(t, "SESSION", sessionCookie.Name)
		assert.NotEmpty(t, sessionCookie.Value)

		sessionID := pgtype.UUID{}
		err = sessionID.Scan(sessionCookie.Value)
		assert.NoError(t, err)
		_, err = queries.GetSession(ctx, sessionID)

		assert.NoError(t, err)
	}
}

func TestSessionCreateShouldFail(t *testing.T) {
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

	UserBuilder := builders.NewUserBuilder(queries).Username("test").Password("test")
	UserBuilder.CreateOne()

	createSessionParams := session.CreateParams{
		Username: "test",
	}

	body, err := json.Marshal(createSessionParams)
	require.NoError(t, err)

	apiClient := utils.NewLoggedInAPIClient(serverURL, "test", "test")
	response, err := apiClient.Post("/api/session", echo.MIMEApplicationJSON, body)
	require.NoError(t, err)

	require.Equal(t, http.StatusBadRequest, response.StatusCode)
	require.Contains(t, string(utils.ReadRespBody(response)), `Error:Field validation for 'Password' failed on the 'required' tag`)
}
