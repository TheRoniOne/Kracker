package internal

import (
	"fmt"
	"log/slog"
)

func GetInput(prompt string) string {
	fmt.Printf("%s: ", prompt)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		slog.Error("Error reading input",
			"error", err)
	}

	return input
}
