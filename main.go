package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	gpt3 "github.com/PullRequestInc/go-gpt3"
)

func main() {
	apiKey := getApiKey()
	context := context.Background()
	client := gpt3.NewClient(apiKey)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please, ask your question?\n")
	scanner.Scan()
	question := scanner.Text()

	if question != "" {
		askAI(client, context, question)
	}
}

func askAI(client gpt3.Client, ctx context.Context, question string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt:    []string{question},
		MaxTokens: gpt3.IntPtr(4000),
	}, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(13)
	}
}

func getApiKey() string {
	apiKey := os.Getenv("AUTOMATE_CODE_API_KEY")

	if apiKey == "" {
		panic("Missing API key")
	}

	return apiKey
}
