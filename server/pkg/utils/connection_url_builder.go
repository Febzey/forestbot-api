package utils

import (
	"fmt"
	"os"
)

func ConnectionUrlBuilder(n string) (string, error) {
	var url string
	switch n {
	case "server":
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)

	default:
		return "", fmt.Errorf("Invalid connection name: %s", n)
	}

	return url, nil
}
