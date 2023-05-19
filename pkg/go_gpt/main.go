package main

import (
	"bufio"
	"context"
	"fmt"
	"giao/pkg/util"
	"github.com/sashabaranov/go-openai"
	"github.com/unidoc/unipdf/v3/common/license"
	"os"
	"strings"
)

// $env:HTTPS_PROXY="http://127.0.0.1:10809"
func main() {
	license.SetMeteredKey("!")
	env := util.NewEnv()
	token := env.Get("talkSecret")
	client := openai.NewClient(token.(string))
	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Conversation")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
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
			continue
		}

		content := resp.Choices[0].Message.Content
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
		fmt.Println(content)
	}

}
