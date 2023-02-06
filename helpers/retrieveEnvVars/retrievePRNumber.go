package retrieveEnvVars

import (
	"fmt"
	"os"
)

type PullRequestEvent struct {
	Number int `json:"number"`
}

func RetrievePRNumber() (string, error) {
	eventNumber := os.Getenv("PULL_REQUEST_NUMBER")
	if eventNumber == "" {
		return "", fmt.Errorf("PULL_REQUEST_NUMBER environment variable not set")
	}

	return eventNumber, nil
}