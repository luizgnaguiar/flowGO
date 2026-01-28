package processor

import (
	"context"
	"flowGO/internal/prompt"
)

type FlowEngine struct {
	Provider interface {
		Execute(ctx context.Context, prompt string) (string, error)
	}
}

func (e *FlowEngine) Process(ctx context.Context, userInput string) (string, error) {
	// Dados que o template espera (pode ser um struct se tiveres mais campos)
	data := struct {
		Input string
	}{
		Input: userInput,
	}

	// Chama o construtor de prompts usando o ficheiro .tmpl
	fullPrompt, err := prompt.BuildPrompt("architect.tmpl", data)
	if err != nil {
		return "", err
	}

	return e.Provider.Execute(ctx, fullPrompt)
}
