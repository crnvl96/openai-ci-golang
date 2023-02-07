package githubClient

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v50/github"
)

type RetrieveFilesArgs struct {
	GHClient	*github.Client
	GHContext context.Context
	RepositoryOwner string
	RepositoryName string
	CommitSHA *string
}

func RetrieveFiles(args RetrieveFilesArgs) *github.RepositoryCommit {
	files, _, err := args.GHClient.Repositories.GetCommit(
		args.GHContext,
		args.RepositoryOwner,
		args.RepositoryName,
		*args.CommitSHA,
		nil,
	)	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return files
}