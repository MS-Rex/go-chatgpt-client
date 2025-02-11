package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
	"strconv"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		log.Fatal("No API key set")
	}
	client := openai.NewClient(apiKey)
	complete(context.Background(), client, "Can you tell me 5 foods")
}

func makeRequest(question string) openai.ChatCompletionRequest {
	maxToken, _ := strconv.Atoi(os.Getenv("MAX_TOKEN"))
	temperature, _ := strconv.ParseFloat(os.Getenv("TEMPERATURE"), 32)
	questions := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleUser,
			Content: question,
		},
	}

	return openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Messages:    questions,
		MaxTokens:   int(maxToken),
		Temperature: float32(temperature),
	}
}

func complete(ctx context.Context, client *openai.Client, question string) {
	request := makeRequest(question)
	resp, _ := client.CreateChatCompletion(ctx, request)
	fmt.Print(resp.Choices[0].Message.Content)
}
