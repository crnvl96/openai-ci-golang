package chatGPTClient

import (
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func GenerateRequestToGPT(content string) gogpt.CompletionRequest {
	request := gogpt.CompletionRequest{
		Model:     gogpt.CodexCodeDavinci002,
		MaxTokens: 2048,
		Temperature: 0.5,
		Prompt: fmt.Sprintf("Explain this code:\n%v\n", content),
	}

	return request
}