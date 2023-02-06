package retrieveEnvVars

import (
	"fmt"
	"os"
)

func RetrieveEnvVars() (string, string, int, error) {
	apiKey, err := RetrieveAIAPIKey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	githubToken, err := RetrieveGithubToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pullRequestNumber, err := RetrievePRNumber()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return apiKey, githubToken, pullRequestNumber, nil
}