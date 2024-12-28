package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/TheRoniOne/Kracker/api/models"
	"github.com/labstack/echo/v4"
)

type LoggedInClient struct {
	Client *http.Client
}

func NewLoggedInClient(url, username, password string) *LoggedInClient {
	client := &http.Client{}

	client.Get(url)

	createSessionParams := models.SessionCreateParams{
		Username: username,
		Password: password,
	}

	j, _ := json.Marshal(createSessionParams)

	response, err := client.Post(url+"/api/session", echo.MIMEApplicationJSON, strings.NewReader(string(j)))
	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusOK {
		panic("Failed to log in")
	}

	return &LoggedInClient{
		Client: client,
	}
}
