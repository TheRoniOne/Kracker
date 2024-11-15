package internal

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_NAME")

	DBUser     = os.Getenv("DB_USER")
	DBPassword = getSecret("DB_PASSWORD")

	Debug = os.Getenv("DEBUG") == "true"

	TimeLocation = getTimeLocation("America/Lima")

	RateLimit = parseIntEnv("RATE_LIMIT")
)

func getSecret(key string) string {
	path := filepath.Join("/run/secrets", strings.ToLower(key))

	contents, err := os.ReadFile(path)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to read secret file: %s", path),
			"error", err)
		slog.Info(fmt.Sprintf("Using env variable %s: as secret", key))

		return os.Getenv(key)
	}

	return strings.TrimSpace(string(contents))
}

func parseIntEnv(key string) int {
	value := os.Getenv(key)

	if value == "" {
		return 0
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to parse %s", key),
			"error", err)
		panic(err)
	}

	return result
}

func getTimeLocation(key string) *time.Location {
	value := os.Getenv(key)

	loc, err := time.LoadLocation(value)
	if err != nil {
		slog.Error("Failed to load time location",
			"error", err)
		panic(err)
	}

	return loc
}
