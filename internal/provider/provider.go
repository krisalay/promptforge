package provider

import (
	"context"

	"github.com/krisalay/promptforge/internal/domain"
)

type ProviderType string

const (
	ProviderTypeOpenAI = "openai"
)

type ChatRequest struct {
	Model        domain.Model
	SystemPrompt string
	UserPrompt   string
	Config       map[string]any
}

type Provider interface {
	Chat(ctx context.Context, req ChatRequest) (string, error)
}
