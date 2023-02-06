package retrieveEnvVars

import (
	"fmt"
	"os"
)

func RetrieveRepoName() (string, error) {
	name := os.Getenv("PULL_REQUEST_NAME")
	if name == "" {
		return "", fmt.Errorf("PULL_REQUEST_NAME environment variable not set")
	}

	return name, nil
}