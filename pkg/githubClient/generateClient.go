package githubClient

import (
	"context"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

func GenerateNewClient(ghToken string) (*github.Client, context.Context) {
	ghContext := context.Background();
	
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghToken},
	)

	tokenClient := oauth2.NewClient(ghContext, tokenSource)
	ghClient := github.NewClient(tokenClient)

	return ghClient, ghContext
}