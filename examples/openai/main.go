package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/krisalay/promptforge"
	"github.com/krisalay/promptforge/internal/domain"
	"github.com/krisalay/promptforge/internal/provider/openai"
)

type Database struct {
	store map[string]*domain.Prompt
	mu    sync.RWMutex
}

func NewDatabase() *Database {
	return &Database{
		store: make(map[string]*domain.Prompt),
	}
}

func (d *Database) SavePrompt(ctx context.Context, prompt *domain.Prompt) error {
	d.store[prompt.Key] = prompt
	return nil
}

func (d *Database) GetPrompt(ctx context.Context, id string) (*domain.Prompt, error) {
	return d.store[id], nil
}

func main() {
	database := NewDatabase()
	provider := openai.NewOpenAIClient("<YOUR_OPENAI_API_KEY>")
	promptForge := promptforge.New(database, provider)
	_ = promptForge.SavePrompt(context.Background(), &domain.Prompt{
		ID:          uuid.New(),
		Key:         "test_prompt",
		Name:        "Test Prompt",
		Description: "Test Prompt Description",
		Template:    "Hello {{.Name}}, you are {{.Age}} years old.",
		Variables: map[string]string{
			"Name": "Name",
			"Age":  "Age",
		},
		Model: "gpt-4o-mini",
		Config: map[string]any{
			"temperature": 0.7,
		},
	})
	renderedPrompt, _ := promptForge.Render(context.Background(), "test_prompt", map[string]string{
		"Name": "Krisalay",
		"Age":  "30",
	})
	res, err := promptForge.Forge(context.Background(), renderedPrompt, "What is the age? Respond in json format.")
	fmt.Println(res, err)
}
