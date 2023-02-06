package githubClient

import (
	"context"
	"fmt"
	"os"

	r "github.com/crnvl96/openai-ci-golang/helpers/retrieveEnvVars"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func GenerateNewClient() (*github.Client, context.Context) {
	token, err := r.RetrieveGithubToken()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	context := context.Background();
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tokenClient := oauth2.NewClient(context, tokenSource)

	client := github.NewClient(tokenClient)

	return client, context
}