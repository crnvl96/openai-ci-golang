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
	FileName        string
}

func RetrieveFileContent(args RetrieveFileContentArgs) string {
	content, _, _, err := args.GHClient.Repositories.GetContents(
		args.GHContext,
		args.RepositoryOwner,
		args.RepositoryName,
		args.FileName,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fileContent, error := content.GetContent()
	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}

	return fileContent
}
