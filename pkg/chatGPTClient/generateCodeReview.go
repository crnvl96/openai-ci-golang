package chatGPTClient

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v50/github"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func GenerateCodeReview(commits []*github.RepositoryCommit, client *github.Client, context context.Context, gptClient *gogpt.Client, gptContext context.Context, owner string, repo string) {
	for _, commit := range commits {
		files, _, err := client.Repositories.GetCommit(context, owner, repo, *commit.SHA, nil)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		

		for _, file := range files.Files {
			fileName := file.Filename
			content, _, _, err := client.Repositories.GetContents(context, owner, repo, *fileName, nil)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			request := gogpt.CompletionRequest{
				Model:     gogpt.CodexCodeDavinci002,
				MaxTokens: 2048,
				Prompt:  "Explain and make suggestions about this code:\n" + *content.Content + "\n",
			}

			response, err := gptClient.CreateCompletion(gptContext, request)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			pullRequest := &github.NewPullRequest{
				Title: fileName,
				Body:  &response.Choices[0].Text,
			}

			_, _, error := client.PullRequests.Create(context, owner, repo, pullRequest)
			if error != nil {
				fmt.Println(error)
				os.Exit(1)
			}
		}
	}
}