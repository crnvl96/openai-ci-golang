package retrieveEnvVars

import (
	"fmt"
	"os"
	"strings"
)

func RetrieveRepoName() (string, error) {
	info := os.Getenv("PULL_REQUEST_INFO")
	if info == "" {
		return "", fmt.Errorf("PULL_REQUEST_INFO environment variable not set")
	}

	name := strings.Split(info, "/")[1]

	return name, nil
}