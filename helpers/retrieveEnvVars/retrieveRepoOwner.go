package retrieveEnvVars

import (
	"fmt"
	"os"
)

func RetrieveRepoOwner() (string, error) {
	owner := os.Getenv("PULL_REQUEST_OWNER")
	if owner == "" {
		return "", fmt.Errorf("PULL_REQUEST_OWNER environment variable not set")
	}

	return owner, nil
}