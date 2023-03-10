package githubClient

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v50/github"
)

type RetrieveCommitsArgs struct{
	PullRequestNumber int
	GHClient *github.Client
	GHContext context.Context
	RepositoryOwner string
	RepositoryName string
}

func RetrieveCommits(args RetrieveCommitsArgs) []*github.RepositoryCommit {
	commits, _, err := args.GHClient.PullRequests.ListCommits(
		args.GHContext,
		args.RepositoryOwner,
		args.RepositoryName,
		args.PullRequestNumber,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return commits
}