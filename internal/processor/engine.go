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
	// CAMADA 1: AUTHOR (Gera o rascunho V1)
	dataV1 := struct{ Input string }{Input: userInput}
	reqV1, _ := prompt.BuildPrompt("architect.tmpl", dataV1)
	v1Output, _ := e.Provider.Execute(ctx, reqV1)

	// CAMADA 2: RUNTIME SIMULATOR (Analisa o V1 e gera os riscos)
	dataSim := struct{ PromptV1 string }{PromptV1: v1Output}
	reqSim, _ := prompt.BuildPrompt("runtime_simulator.tmpl", dataSim)
	risksOutput, _ := e.Provider.Execute(ctx, reqSim)

	// CAMADA 3: CONSTRAINT ENFORCER (Une V1 + Riscos para gerar V2)
	dataV2 := struct {
		PromptV1     string
		RuntimeRisks string
	}{
		PromptV1:     v1Output,
		RuntimeRisks: risksOutput, // Agora os riscos vêm da IA!
	}
	reqV2, _ := prompt.BuildPrompt("constraint_enforcer.tmpl", dataV2)

	return e.Provider.Execute(ctx, reqV2)
}
