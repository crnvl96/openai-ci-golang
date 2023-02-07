package main

import (
	"context"

	a "github.com/crnvl96/openai-ci-golang/helpers/args"
	gpt "github.com/crnvl96/openai-ci-golang/pkg/chatGPTClient"
	gh "github.com/crnvl96/openai-ci-golang/pkg/githubClient"
	"github.com/google/go-github/v50/github"
	gogpt "github.com/sashabaranov/go-gpt3"
)

var args a.Args
var gptClient *gogpt.Client
var gptContext context.Context
var ghClient *github.Client
var ghContext context.Context


func init() {
	args = a.GetAllArgs()
	gptClient, gptContext = gpt.GenerateNewClient(args.AIAPIKey)
	ghClient, ghContext = gh.GenerateNewClient(args.GithubToken)
}


func main() {
	gpt.GenerateCodeReview(
		gpt.GenerateCodeReviewArgs{
			GHClient: ghClient,
			GHContext: ghContext,
			GPTClient: gptClient,
			GPTContext: gptContext,
			RepositoryOwner: args.RepositoryOwner,
			RepositoryName: args.RepositoryName,
			PullRequestNumber: args.PullRequestNumber,
		},
	)
}