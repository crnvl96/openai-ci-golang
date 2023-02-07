package chatGPTClient

import (
	"context"
	"fmt"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type GetCompletionArgs struct {
	FileName string
	GPTClient *gogpt.Client
	GPTContext context.Context
	request gogpt.CompletionRequest
}

func GetCompletion(args GetCompletionArgs) string {
	completion, err := args.GPTClient.CreateCompletion(args.GPTContext, args.request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	text := completion.Choices[0].Text
	response := fmt.Sprintf("%v\n\n%v\n", args.FileName, text)

	return response
}