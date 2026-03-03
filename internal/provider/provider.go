package provider

import (
	"context"
)

type ProviderType string

const (
	ProviderTypeOpenAI = "openai"
)

type ChatRequest struct {
	Model        string
	SystemPrompt string
	UserPrompt   string
	Config       map[string]any
}

type Provider interface {
	Chat(ctx context.Context, req ChatRequest) (string, error)
}
