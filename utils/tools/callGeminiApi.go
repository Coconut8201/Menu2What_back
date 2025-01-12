package tools

import (
	"Menu2What_back/utils/Logger"
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
)

// GetGeminiModel returns a generative model instance
func GetGeminiModel(client *genai.Client) *genai.GenerativeModel {
	return client.GenerativeModel("gemini-1.5-flash")
}

func CallGeminiApi(client *genai.Client, prompt string) (string, error) {
	logger, err := Logger.NewLogger(Logger.INFO)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	model := GetGeminiModel(client)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		logger.Info("Failed to generate content: " + err.Error())
		return "", err
	}

	fmt.Println(resp)
	// return resp.Text(), nil
	return "", nil
}
