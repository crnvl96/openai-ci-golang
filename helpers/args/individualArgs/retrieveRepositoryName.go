package individualArgs

import (
	"fmt"
	"os"
	"strings"
)

func RetrieveRepositoryName() (string, error) {
	info := os.Getenv("PULL_REQUEST_INFO")	
	if info == "" {
		return "", fmt.Errorf("PULL_REQUEST_INFO environment variable not set")
	}

	repoName := strings.Split(info, "/")[1]

	return repoName, nil
}