package individualArgs

import (
	"fmt"
	"os"
)

func RetrieveAIAPIKey() (string, error) {
	APIKey := os.Getenv("API_KEY")	
	if APIKey == "" {
		return "", fmt.Errorf("API_KEY environment variable not set")
	}
	
	return APIKey, nil
}