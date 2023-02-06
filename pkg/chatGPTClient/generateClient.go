package chatGPTClient

import (
	"context"
	"fmt"
	"os"

	r "github.com/crnvl96/openai-ci-golang/helpers/retrieveEnvVars"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func GenerateClient() (*gogpt.Client, context.Context) {
	openaiKey, err := r.RetrieveAIAPIKey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gptClient := gogpt.NewClient(openaiKey)
	context := context.Background()

	return gptClient, context
}