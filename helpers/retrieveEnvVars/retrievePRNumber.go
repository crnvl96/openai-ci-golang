package retrieveEnvVars

import (
	"encoding/json"
	"fmt"
	"os"
)

type PullRequestEvent struct {
	Number int `json:"number"`
}

func RetrievePRNumber() (int, error) {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	if eventPath == "" {
		return 0, fmt.Errorf("GITHUB_EVENT_PATH environment variable not set")
	}

	file, err := os.Open(eventPath)
	if err != nil {
		return 0, fmt.Errorf("error opening GITHUB_EVENT_PATH file: %v", err)
	}
	defer file.Close()

	var event PullRequestEvent
	err = json.NewDecoder(file).Decode(&event)
	if err != nil {
		return 0, fmt.Errorf("error decoding pull request event: %v", err)
	}

	return event.Number, nil
}