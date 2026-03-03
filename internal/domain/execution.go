package domain

import "github.com/google/uuid"

type Model string

const (
	ModelGPT4oMini Model = "gpt-4o-mini"
)

type Execution struct {
	ID             uuid.UUID
	PromptID       uuid.UUID
	PromptVersion  int
	InputVariables map[string]string
	Output         string
	TokensUsed     int
	LatencyMS      int
	Status         string
	Model          Model
}
