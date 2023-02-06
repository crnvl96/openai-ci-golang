package retrieveEnvVars

import (
	"fmt"
	"os"
)

func RetrieveAIAPIKey() (string, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("API_KEY environment variable not set")
	}
	
	return apiKey, nil
}