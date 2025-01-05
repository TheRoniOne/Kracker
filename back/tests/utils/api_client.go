package utils

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/TheRoniOne/Kracker/api/models/session"
	"github.com/labstack/echo/v4"
)

type APIClient struct {
	Client  *http.Client
	BaseURL string
}

func NewAPIClient(baseURL string) *APIClient {
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

func NewLoggedInAPIClient(baseURL, username, password string) *APIClient {
	client := NewAPIClient(baseURL)

	createSessionParams := session.CreateParams{
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

func (c *APIClient) Post(path string, contentType string, body []byte) (*http.Response, error) {
	return c.Client.Post(c.MakeURL(path), contentType, strings.NewReader(string(body)))
}

func (c *APIClient) Patch(path string, contentType string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("PATCH", c.MakeURL(path), bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", contentType)

	return c.Client.Do(req)
}
