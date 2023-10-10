package prompt

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func Index(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"apiName": "Yoda AI",
		"status":  "running",
	})
}

func Prompt(c *gin.Context) {
	apiKey, ok := viper.Get("OPENAPI_KEY").(string)

	// get the prompt
	prompt := c.DefaultQuery("prompt", "")

	if prompt == "" || !ok {

		c.JSON(http.StatusBadRequest, gin.H{"error": "No prompt provided."})
	}

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: 800,
		},
	)

	if err != nil {

		fmt.Printf("ChatCompletion error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Prompt error. Holmes AI is out of tokens."})
	}

	fmt.Println(resp.Choices[0].Message.Content)

	c.JSON(http.StatusOK, gin.H{
		"prompt": prompt,
		"reply":  resp.Choices[0].Message.Content,
	})
}
