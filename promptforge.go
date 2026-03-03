package promptforge

import (
	"bytes"
	"context"
	"text/template"

	"github.com/krisalay/promptforge/internal/domain"
	"github.com/krisalay/promptforge/internal/provider"
	"github.com/krisalay/promptforge/pkg/memcache"
)

type PromptForge struct {
	storage  Storage
	provider provider.Provider
	memCache memcache.MemCache[*domain.PromptCache]
}

func New(storage Storage, provider provider.Provider) *PromptForge {
	return &PromptForge{
		storage:  storage,
		provider: provider,
		memCache: memcache.NewMemCache[*domain.PromptCache](),
	}
}

func (p *PromptForge) SavePrompt(ctx context.Context, prompt *domain.Prompt) error {
	if err := p.storage.SavePrompt(ctx, prompt); err != nil {
		return err
	}
	return p.memCache.Remove(ctx, prompt.ID.String())
}

func (p *PromptForge) Render(ctx context.Context, id string, vars map[string]string) (*domain.RenderedPrompt, error) {
	promptCache, err := p.memCache.Get(ctx, id)
	if err == nil {
		return p.parseTmpl(ctx, promptCache, vars)
	}

	prompt, err := p.storage.GetPrompt(ctx, id)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New(id).Parse(prompt.Template)
	if err != nil {
		return nil, err
	}
	promptCache.Tmpl = tmpl

	return p.parseTmpl(ctx, promptCache, vars)
}

func (p *PromptForge) Forge(ctx context.Context, prompt *domain.RenderedPrompt, userPrompt string) (string, error) {
	return p.provider.Chat(ctx, provider.ChatRequest{
		Model:        prompt.Model,
		SystemPrompt: prompt.Prompt,
		UserPrompt:   userPrompt,
		Config:       prompt.Config,
	})
}

func (p *PromptForge) parseTmpl(_ context.Context, promptCache *domain.PromptCache, vars map[string]string) (*domain.RenderedPrompt, error) {
	var buf bytes.Buffer
	if err := promptCache.Tmpl.Execute(&buf, vars); err != nil {
		return nil, err
	}
	return &domain.RenderedPrompt{
		Model:  promptCache.Model,
		Prompt: buf.String(),
		Config: promptCache.Config,
	}, nil
}
