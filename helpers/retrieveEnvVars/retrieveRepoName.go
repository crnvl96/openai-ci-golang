package retrieveEnvVars

import (
	"fmt"
	"os"
)

func RetrieveRepoName() (string, error) {
	info := os.Getenv("PULL_REQUEST_INFO")
	if info == "" {
		return "", fmt.Errorf("PULL_REQUEST_INFO environment variable not set")
	}

	return "teste", nil
}