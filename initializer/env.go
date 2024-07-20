package initializer

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	env := os.Getenv("KRACKER_ENV")

	if env == "" {
		env = ".local"
	}

	if err := godotenv.Load(".env" + env); err != nil {
		panic(err)
	}
}
