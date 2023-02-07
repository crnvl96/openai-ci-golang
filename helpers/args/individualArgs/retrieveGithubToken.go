package individualArgs

import (
	"fmt"
	"os"
)

func RetrieveGithubToken() (string, error) {
	githubToken := os.Getenv("GITHUB_TOKEN")	
	if githubToken == "" {
		return "", fmt.Errorf("GITHUB_TOKEN environment variable not set")
	}
	
	return githubToken, nil
}