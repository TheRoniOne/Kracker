package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/TheRoniOne/Kracker/api"
	"github.com/TheRoniOne/Kracker/api/models"
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

	userData := models.UserCreateParams{
		Username:  "test",
		Email:     "test",
		Password:  "test",
		Firstname: "test",
		Lastname:  "test",
	}

	j, _ := json.Marshal(userData)

	_, err = http.Post(serverURL+"/api/user/create", echo.MIMEApplicationJSON, strings.NewReader(string(j)))
	require.NoError(t, err)

	createSessionParams := models.SessionCreateParams{
		Username: userData.Username,
		Password: userData.Password,
	}

	j, _ = json.Marshal(createSessionParams)

	response, err := http.Post(serverURL+"/api/session", echo.MIMEApplicationJSON, strings.NewReader(string(j)))
	require.NoError(t, err)

	if assert.Equal(t, http.StatusOK, response.StatusCode) {
		sessionCookie := response.Cookies()[0]

		assert.Equal(t, "SESSION", sessionCookie.Name)
		assert.NotEmpty(t, sessionCookie.Value)

		sessionID := pgtype.UUID{}
		err = sessionID.Scan(sessionCookie.Value)
		assert.NoError(t, err)
		_, err = queries.GetSession(ctx, sessionID)

		assert.NoError(t, err)
	}
}
