package internal

import (
	"fmt"
	"log/slog"
	"os"
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

	JWTSecret   = getSecret("JWT_SECRET")
	JWTLifetime = time.Hour * time.Duration(parseIntEnv("JWT_LIFETIME_HOURS"))

	RateLimit = parseIntEnv("RATE_LIMIT")
)

func getSecret(key string) string {
	path := os.Getenv(key)

	if path == "" {
		return ""
	}

	contents, err := os.ReadFile(path)
	if err != nil {
		logger := slog.Default()

		logger.Error(fmt.Sprintf("Failed to read secret file: %s", path))

		envVar, _ := strings.CutSuffix(key, "_SECRET")
		logger.Info(fmt.Sprintf("Using env variable %s: as secret", envVar))

		return os.Getenv(envVar)
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
		logger := slog.Default()

		logger.Error(fmt.Sprintf("Failed to parse %s: %v", key, err))
		panic(err)
	}

	return result
}
