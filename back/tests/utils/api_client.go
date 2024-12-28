package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/TheRoniOne/Kracker/api/models"
	"github.com/labstack/echo/v4"
)

type APIClient struct {
	Client  *http.Client
	BaseURL string
}

func NewClient(baseURL string) *APIClient {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Get(baseURL)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("Failed to connect to Kracker API")
	}

	return &APIClient{
		Client:  client,
		BaseURL: baseURL,
	}
}

func NewLoggedInClient(baseURL, username, password string) *APIClient {
	client := NewClient(baseURL)

	createSessionParams := models.SessionCreateParams{
		Username: username,
		Password: password,
	}

	j, _ := json.Marshal(createSessionParams)

	resp, err := client.Client.Post(client.MakeURL("/api/session"), echo.MIMEApplicationJSON, strings.NewReader(string(j)))
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		slog.Error("Failed to create session",
			"status", resp.Status,
			"body", ReadRespBody(resp))
		panic("Failed to log in")
	}

	return client
}

func (c *APIClient) MakeURL(path string) string {
	return c.BaseURL + path
}

func (c *APIClient) Get(path string) (*http.Response, error) {
	return c.Client.Get(c.MakeURL(path))
}
