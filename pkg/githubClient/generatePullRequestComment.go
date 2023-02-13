package githubClient

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v50/github"
)

type GeneratePullRequestCommentArgs struct {
	Body              string
	GHClient          *github.Client
	GHContext         context.Context
	RepositoryOwner   string
	RepositoryName    string
	PullRequestNumber int
}

func GeneratePullRequestComment(args GeneratePullRequestCommentArgs) {
	newComment := &github.IssueComment{
		Body: &args.Body,
	}

	_, _, error := args.GHClient.Issues.CreateComment(
		args.GHContext,
		args.RepositoryOwner,
		args.RepositoryName,
		args.PullRequestNumber,
		newComment,
	)

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
}
