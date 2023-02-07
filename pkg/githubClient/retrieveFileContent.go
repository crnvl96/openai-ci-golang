package githubClient

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v50/github"
)

type RetrieveFileContentArgs struct {
	GHClient        *github.Client
	GHContext       context.Context
	RepositoryOwner string
	RepositoryName  string
	FilePath        string
	Branch          string
}

func RetrieveFileContent(args RetrieveFileContentArgs) string {
	content, _, _, err := args.GHClient.Repositories.GetContents(
		args.GHContext,
		args.RepositoryOwner,
		args.RepositoryName,
		args.FilePath,
		&github.RepositoryContentGetOptions{Ref: args.Branch},
	)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fileContent, err := content.GetContent()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return fileContent
}
