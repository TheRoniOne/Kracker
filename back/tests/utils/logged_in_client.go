package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/TheRoniOne/Kracker/api/models"
	"github.com/labstack/echo/v4"
)

type LoggedInClient struct {
	Client  *http.Client
	BaseURL string
}

func NewLoggedInClient(baseURL, username, password string) *LoggedInClient {
	client := &http.Client{}

	client.Get(baseURL)

	createSessionParams := models.SessionCreateParams{
		Username: username,
		Password: password,
	}

	j, _ := json.Marshal(createSessionParams)

	response, err := client.Post(baseURL+"/api/session", echo.MIMEApplicationJSON, strings.NewReader(string(j)))
	if err != nil {
		panic(err)
	}

	if response.StatusCode != http.StatusOK {
		slog.Error(response.Status)
		panic("Failed to log in")
	}

	return &LoggedInClient{
		Client:  client,
		BaseURL: baseURL,
	}
}

func (c *LoggedInClient) PathEncode(path string) string {
	return c.BaseURL + url.QueryEscape(path)
}

func (c *LoggedInClient) Get(path string) *http.Response {
	req, err := http.NewRequest(http.MethodGet, c.PathEncode(path), nil)
	if err != nil {
		panic(err)
	}

	response, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}

	return response
}

func (c *LoggedInClient) Post(path string, contentType string, body *strings.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, c.PathEncode(path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	return c.Client.Do(req)
}

func (c *LoggedInClient) Patch(path string, contentType string, body *strings.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, c.PathEncode(path), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)

	return c.Client.Do(req)
}

func (c *LoggedInClient) Delete(path string) *http.Response {
	req, err := http.NewRequest(http.MethodDelete, c.PathEncode(path), nil)
	if err != nil {
		panic(err)
	}

	response, err := c.Client.Do(req)
	if err != nil {
		panic(err)
	}

	return response
}
