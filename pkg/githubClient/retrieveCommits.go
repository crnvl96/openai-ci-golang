package githubClient

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/v50/github"
)

func RetrieveCommits(client *github.Client, context context.Context, owner string, repo string, pullRequestId string) ([]*github.RepositoryCommit, error) {
	pullRequestNumber, err := strconv.Atoi(pullRequestId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	commits, _, err := client.PullRequests.ListCommits(context, owner, repo, pullRequestNumber, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return commits, nil	
}