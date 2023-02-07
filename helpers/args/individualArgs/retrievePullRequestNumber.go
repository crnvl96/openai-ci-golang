package individualArgs

import (
	"fmt"
	"os"
	"strconv"
)

func RetrievePullRequestNumber() (int, error) {
	pullRequestId := os.Getenv("PULL_REQUEST_NUMBER")
	if pullRequestId == "" {
		return 0, fmt.Errorf("PULL_REQUEST_NUMBER environment variable not set")
	}

	pullRequestNumber, err := strconv.Atoi(pullRequestId)
	if err != nil {
		return 0, fmt.Errorf("PULL_REQUEST_NUMBER environment variable is not an integer")
	}

	return pullRequestNumber, nil
}