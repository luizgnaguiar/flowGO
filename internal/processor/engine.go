package processor

import (
	"context"
	"fmt"
)

type FlowEngine struct {
	Provider interface {
		Execute(ctx context.Context, prompt string) (string, error)
	}
}

func (e *FlowEngine) Process(ctx context.Context, userInput string) (string, error) {
	// Aqui defines os teus PARÂMETROS de sistema
	systemPrompt := "Atua como um arquiteto de software sénior. " +
		"Analisa o seguinte pedido e gera um plano técnico: "

	fullPrompt := fmt.Sprintf("%s\n\nUser Input: %s", systemPrompt, userInput)

	return e.Provider.Execute(ctx, fullPrompt)
}
