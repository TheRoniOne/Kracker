package internal

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
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

	Debug = parseBoolEnv("DEBUG", false)

	TimeLocation = getTimeLocation("TIME_LOCATION")

	RateLimit = parseIntEnv("RATE_LIMIT", 10)

	RootPath          = getRootPath()
	SessionMaxAgeDays = parseIntEnv("SESSION_MAX_AGE_DAYS", 30)

	DOMAIN           = os.Getenv("DOMAIN")
	CSRFCookieSecure = parseBoolEnv("CSRF_COOKIE_SECURE", false)
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

func parseIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to parse %s", key),
			"error", err)
		panic(err)
	}

	return result
}

func parseBoolEnv(key string, defaultValue bool) bool {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	result, err := strconv.ParseBool(value)
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

func getRootPath() string {
	_, b, _, _ := runtime.Caller(0)

	return filepath.Join(filepath.Dir(b), "../")
}
