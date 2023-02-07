package chatGPTClient

import (
	gogpt "github.com/sashabaranov/go-gpt3"
)

func GenerateRequestToGPT(content string) gogpt.CompletionRequest {
	t := "Explain code: func Diff(a, b int) int {\n\treturn a - b\n}"

	request := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci002,
		MaxTokens:   2048,
		Temperature: 0.5,
		Prompt:      t,
		// Prompt:      fmt.Sprintf("Explain Code: \n %s", content),
	}

	return request
}
