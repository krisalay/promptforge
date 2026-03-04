package promptforge

import (
	"context"

	"github.com/krisalay/promptforge/internal/domain"
)

type Storage interface {
	SavePrompt(ctx context.Context, prompt *domain.Prompt) error
	GetPrompt(ctx context.Context, id string) (*domain.Prompt, error)
}
