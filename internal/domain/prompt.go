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
	Model       Model
	Config      map[string]any
}

type PromptCache struct {
	Model  Model
	Tmpl   *template.Template
	Config map[string]any
}

type RenderedPrompt struct {
	Model  Model
	Prompt string
	Config map[string]any
}
