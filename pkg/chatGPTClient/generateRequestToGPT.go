package chatGPTClient

import (
	"fmt"

	gogpt "github.com/sashabaranov/go-gpt3"
)

func GenerateRequestToGPT(content string) gogpt.CompletionRequest {
	request := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci002,
		MaxTokens:   2047,
		Temperature: 0.5,
		Prompt:      fmt.Sprintf("Explain Code: \n %s", content),
	}

	return request
}
