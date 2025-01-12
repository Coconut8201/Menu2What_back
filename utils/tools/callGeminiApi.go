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

// CallGeminiApiTextOnly 僅輸入文字 並調用 Gemini API 並返回回應
func CallGeminiApiTextOnly(prompt string) (string, error) {
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
	//D printResponse(resp)
	return resultString, nil
}

func CallGeminiApiImageAndText(prompt string, imageData []byte) (string, error) {
	logger, err := Logger.NewLogger(Logger.INFO)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("Gemini_API_Key")))
	if err != nil {
		logger.Info("Failed to create client: " + err.Error())
		return "", err
	}
	defer client.Close()

	model := GetGeminiModel(client)
	resp, err := model.GenerateContent(ctx,
		genai.Text(prompt),
		genai.ImageData("jpeg", imageData))
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
	//D printResponse(resp)
	return resultString, nil
}
