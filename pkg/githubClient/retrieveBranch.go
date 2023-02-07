package githubClient

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v50/github"
)

type RetrieveBranchArgs struct {
	GHClient          *github.Client
	GHContext         context.Context
	RepositoryOwner   string
	RepositoryName    string
	PullRequestNumber int
}

func RetrieveBranch(args RetrieveBranchArgs) string {
	pullRequest, _, err := args.GHClient.PullRequests.Get(
		args.GHContext,
		args.RepositoryOwner,
		args.RepositoryName,
		args.PullRequestNumber,
	)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	branch := pullRequest.GetHead().GetRef()

	return branch
}
