package tools

import (
	"Menu2What_back/utils/Logger"
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// printResponse 打印 Gemini API 的回應
func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}

// GetGeminiModel returns a generative model instance
func GetGeminiModel(client *genai.Client) *genai.GenerativeModel {
	return client.GenerativeModel("gemini-1.5-flash")
}

func CallGeminiApi(prompt string) (string, error) {
	logger, err := Logger.NewLogger(Logger.INFO)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("Gemini_API_Key")))
	model := GetGeminiModel(client)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		logger.Info("Failed to generate content: " + err.Error())
		return "", err
	}

	var resultString string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				resultString += fmt.Sprintf("%v", part)
			}
		}
	}

	printResponse(resp)

	return resultString, nil
}
