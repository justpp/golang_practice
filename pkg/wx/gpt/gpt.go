package gpt

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

var (
	msgCache map[string][]openai.ChatCompletionMessage
)

func init() {
	msgCache = make(map[string][]openai.ChatCompletionMessage)
}

func Talk(token, fromUser, question string) string {
	client := openai.NewClient(token)
	messages := msgCache[fromUser]

	newTalk := false
	if len(messages) > 20 {
		messages = make([]openai.ChatCompletionMessage, 0)
		newTalk = true
	}
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: question,
	})
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}

	content := resp.Choices[0].Message.Content
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	fmt.Println("messages", messages)
	msgCache[fromUser] = messages
	if newTalk {
		return content + " (7秒钟之后的鱼说)"
	}
	return content
}
