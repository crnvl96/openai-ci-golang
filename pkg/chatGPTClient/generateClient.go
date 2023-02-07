package chatGPTClient

import (
	"context"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func GenerateNewClient(AIAPIKey string) (*gogpt.Client, context.Context) {
	gptClient := gogpt.NewClient(AIAPIKey)
	gptContext := context.Background()

	return gptClient, gptContext
}