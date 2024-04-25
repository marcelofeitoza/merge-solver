package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type MergeRequest struct {
	New      string `json:"new"`
	Old      string `json:"old"`
	Rejected string `json:"rejected,omitempty"`
}

type OpenAI struct {
	client  *openai.Client
	context context.Context
}

func NewOpenAI() *OpenAI {
	secretKey := os.Getenv("OPENAI_SECRET_KEY")
	fmt.Println("secretKey: ", secretKey)
	if secretKey == "" {
		log.Fatalf("OPENAI_SECRET_KEY is required")
	}

	client := openai.NewClient(secretKey)
	return &OpenAI{
		client:  client,
		context: context.Background(),
	}
}

func (o *OpenAI) Merge(request MergeRequest) (openai.ChatCompletionMessage, error) {
	message, err := json.Marshal(request)
	if err != nil {
		return openai.ChatCompletionMessage{}, err
	}

	resp, err := o.client.CreateChatCompletion(
		o.context,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo16K,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are an advanced code merge tool. You will receive two sets of Rust code: 'old' and 'new'. Your task is to intelligently integrate the 'new' code into the 'old', prioritizing the changes from the 'new'. Direct conflicts should be resolved in favor of the 'new' code. Any part of the 'old' code that is unchanged or does not conflict with the 'new' code should remain intact. Deprecations and removals in the 'new' code should be respected. Output the merged code that compiles successfully and is ready for developer review, emphasizing functionality, cleanliness, and the prevention of code regression. Also, if it happens to have a `rejected` field, it means that the code was not accepted by the merge tool, thus, it should be regenerated.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: string(message),
				},
			},
		},
	)

	if err != nil {
		// return openai.ChatCompletionMessage{}, err
		return openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleSystem,
			Content: "I'm sorry, I couldn't merge the code. Please try again.",
		}, nil
	}

	return resp.Choices[0].Message, nil
}
