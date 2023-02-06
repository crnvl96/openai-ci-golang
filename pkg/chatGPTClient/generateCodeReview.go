package chatGPTClient

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/v50/github"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func GenerateCodeReview(commits []*github.RepositoryCommit, client *github.Client, context context.Context, gptClient *gogpt.Client, gptContext context.Context, owner string, repo string, pullRequest string) {
	prNumber, err := strconv.Atoi(pullRequest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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

			body := fmt.Sprintf("%v\n\n%v\n", *fileName, response.Choices[0].Text)

			comment := &github.IssueComment{
				Body: &body,
			}

			_, _, error := client.Issues.CreateComment(context, owner, repo, prNumber, comment)
			if error != nil {
				fmt.Println("Error:", error)
				return
			}
		}
	}
}