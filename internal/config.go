package internal

import (
	"os"
	"strings"

	"github.com/labstack/gommon/log"
)

var (
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")

	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = getSecret("DB_PASSWORD")

	DEBUG = os.Getenv("DEBUG") == "true"
)

func getSecret(key string) string {
	path := os.Getenv(key)

	if path == "" {
		return ""
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		log.Errorf("Failed to read secret file: %s", path)

		env_var, _ := strings.CutSuffix(key, "_SECRET")
		log.Infof("Using env variable: %s as secret", env_var)

		return os.Getenv(env_var)
	}

	return strings.TrimSpace(string(contents))
}
