package chatGPTClient

import (
	"context"
	"fmt"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type GetCompletionArgs struct {
	GPTClient  *gogpt.Client
	GPTContext context.Context
	request    gogpt.CompletionRequest
}

func GetCompletion(args GetCompletionArgs) string {
	completion, err := args.GPTClient.CreateCompletion(args.GPTContext, args.request)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("test")

	return completion.Choices[0].Text
}
