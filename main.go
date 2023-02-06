package main

import (
	"fmt"
	"os"

	r "github.com/crnvl96/openai-ci-golang/helpers/retrieveEnvVars"
	"github.com/crnvl96/openai-ci-golang/pkg/chatGPTClient"
	"github.com/crnvl96/openai-ci-golang/pkg/githubClient"
)

func main() {
	gptClient, gptContext := chatGPTClient.GenerateClient()
	GHClient, GHContext := githubClient.GenerateNewClient()

	owner, err := r.RetrieveRepoOwner()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repo, err := r.RetrieveRepoName()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pullRequest, err := r.RetrievePRNumber()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	commits, err := githubClient.RetrieveCommits(GHClient, GHContext, owner, repo, pullRequest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	chatGPTClient.GenerateCodeReview(commits, GHClient, GHContext, gptClient, gptContext, owner, repo)
}