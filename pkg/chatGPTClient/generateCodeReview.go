package chatGPTClient

import (
	"context"

	gh "github.com/crnvl96/openai-ci-golang/pkg/githubClient"
	"github.com/google/go-github/v50/github"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type GenerateCodeReviewArgs struct {
	GHClient          *github.Client
	GHContext         context.Context
	GPTClient         *gogpt.Client
	GPTContext        context.Context
	RepositoryOwner   string
	RepositoryName    string
	PullRequestNumber int
}

func GenerateCodeReview(args GenerateCodeReviewArgs) {
	commits := gh.RetrieveCommits(
		gh.RetrieveCommitsArgs{
			PullRequestNumber: args.PullRequestNumber,
			GHClient:          args.GHClient,
			GHContext:         args.GHContext,
			RepositoryOwner:   args.RepositoryOwner,
			RepositoryName:    args.RepositoryName,
		},
	)

	for _, commit := range commits {
		filesFromCommit := gh.RetrieveFiles(
			gh.RetrieveFilesArgs{
				GHClient:        args.GHClient,
				GHContext:       args.GHContext,
				RepositoryOwner: args.RepositoryOwner,
				RepositoryName:  args.RepositoryName,
				CommitSHA:       commit.SHA,
			},
		)

		for _, file := range filesFromCommit.Files {
			fileName := *file.Filename

			fileContent := gh.RetrieveFileContent(
				gh.RetrieveFileContentArgs{
					GHClient:        args.GHClient,
					GHContext:       args.GHContext,
					RepositoryOwner: args.RepositoryOwner,
					RepositoryName:  args.RepositoryName,
					FileName:        fileName,
				},
			)

			request := GenerateRequestToGPT(fileContent)

			response := GetCompletion(
				GetCompletionArgs{
					FileName:   fileName,
					GPTClient:  args.GPTClient,
					GPTContext: args.GPTContext,
					request:    request,
				},
			)

			gh.GeneratePullRequestComment(
				gh.GeneratePullRequestCommentArgs{
					Body:              response,
					GHClient:          args.GHClient,
					GHContext:         args.GHContext,
					RepositoryOwner:   args.RepositoryOwner,
					RepositoryName:    args.RepositoryName,
					PullRequestNumber: args.PullRequestNumber,
				},
			)
		}
	}
}
