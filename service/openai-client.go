package service

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type OpenAiClientWrapper struct {
	client *openai.Client
}

func NewOpenAiClient(apiKey string) OpenAiClientWrapper {
	fmt.Println(apiKey)
	return OpenAiClientWrapper{
		client: openai.NewClient(apiKey),
	}
}

func (wrapper OpenAiClientWrapper) SendRequest(prompt string) (string, error) {
	resp, err := wrapper.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
