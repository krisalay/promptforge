package domain

import (
	"text/template"

	"github.com/google/uuid"
)

type Prompt struct {
	ID          uuid.UUID
	Key         string
	Name        string
	Description string
	Template    string
	Variables   map[string]string
	Model       string
	Config      map[string]any
}

type PromptCache struct {
	Model  string
	Tmpl   *template.Template
	Config map[string]any
}

type RenderedPrompt struct {
	Model  string
	Prompt string
	Config map[string]any
}
