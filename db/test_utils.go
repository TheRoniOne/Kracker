package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/TheRoniOne/Kracker/internal"
)

var Migrations []string

func init() {
	var err error
	migrationsPath := filepath.Join(internal.RootPath, "db/migrations/*.sql")
	Migrations, err = filepath.Glob(migrationsPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while trying to find migrations: %v\n", err)
		os.Exit(1)
	}
}
