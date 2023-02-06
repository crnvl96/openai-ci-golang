package retrieveEnvVars

import (
	"fmt"
	"os"
	"strings"
)

func RetrieveRepoOwner() (string, error) {
	info := os.Getenv("PULL_REQUEST_INFO")
	if info == "" {
		return "", fmt.Errorf("PULL_REQUEST_INFO environment variable not set")
	}

	owner := strings.Split(info, "/")[0]

	return owner, nil
}