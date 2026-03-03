package promptforge

import (
	"context"

	"github.com/krisalay/promptforge/internal/domain"
)

type Storage interface {
	SavePrompt(ctx context.Context, prompt *domain.Prompt) error
	GetPrompt(ctx context.Context, id string) (*domain.Prompt, error)
}

type Cache interface {
	Get(ctx context.Context, id string) (string, error)
	Set(ctx context.Context, id, parsedPrompt string) error
	Invalidate(ctx context.Context, id string) error
}
