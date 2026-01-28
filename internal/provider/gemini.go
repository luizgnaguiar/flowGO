package provider

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiClient struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func NewGeminiClient(ctx context.Context, apiKey string) (*GeminiClient, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}
	// Usando o modelo que confirmaste que funciona
	model := client.GenerativeModel("gemini-2.0-flash")
	return &GeminiClient{client: client, model: model}, nil
}

func (g *GeminiClient) Execute(ctx context.Context, prompt string) (string, error) {
	resp, err := g.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}
	return string(resp.Candidates[0].Content.Parts[0].(genai.Text)), nil
}
