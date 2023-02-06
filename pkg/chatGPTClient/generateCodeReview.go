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
				Temperature: 0.5,
				Prompt:  "Explain this code:\n" + *content.Content + "\n",
			}

			response, err := gptClient.CreateCompletion(gptContext, request)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			issue := &github.IssueRequest{
				Title: fileName,
				Body:  &response.Choices[0].Text,
			}

			_, _, errorr := client.Issues.Create(context, owner, repo, issue)
			if errorr != nil {
				fmt.Println(errorr)
				os.Exit(1)
			}
		}
	}
}