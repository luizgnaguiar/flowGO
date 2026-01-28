package processor

import (
	"context"
	"flowGO/internal/prompt"
	"fmt"
)

type FlowEngine struct {
	Provider interface {
		Execute(ctx context.Context, prompt string) (string, error)
	}
}

func (e *FlowEngine) Process(ctx context.Context, userInput string) (string, error) {
	// --- CAMADA 1: AUTHOR ---
	dataV1 := struct{ Input string }{Input: userInput}
	reqV1, _ := prompt.BuildPrompt("architect.tmpl", dataV1)
	v1Output, _ := e.Provider.Execute(ctx, reqV1)
	fmt.Println("\n[DEBUG] Camada 1 (Author) concluída.")

	// --- CAMADA 2: RUNTIME SIMULATOR ---
	dataSim := struct{ PromptV1 string }{PromptV1: v1Output}
	reqSim, _ := prompt.BuildPrompt("runtime_simulator.tmpl", dataSim)
	risksOutput, _ := e.Provider.Execute(ctx, reqSim)
	fmt.Println("[DEBUG] Camada 2 (Simulator) identificou riscos.")

	// --- CAMADA 3: CONSTRAINT ENFORCER ---
	dataV2 := struct {
		PromptV1     string
		RuntimeRisks string
	}{
		PromptV1:     v1Output,
		RuntimeRisks: risksOutput,
	}
	reqV2, _ := prompt.BuildPrompt("constraint_enforcer.tmpl", dataV2)
	v2Output, err := e.Provider.Execute(ctx, reqV2)
	if err != nil {
		return "", err
	}
	fmt.Println("[DEBUG] Camada 3 (Enforcer) gerou Prompt_v2.")

	// --- CAMADA 4: CODE REVIEWER (Nova!) ---
	dataRev := struct{ PromptV2 string }{PromptV2: v2Output}
	reqRev, _ := prompt.BuildPrompt("code_reviewer.tmpl", dataRev)
	reviewOutput, err := e.Provider.Execute(ctx, reqRev)
	if err != nil {
		return "", err
	}
	fmt.Println("[DEBUG] Camada 4 (Reviewer) finalizou a avaliação.")

	// Para o resultado final, vamos concatenar o Prompt_v2 com o Review
	// para que o usuário veja o contrato E a revisão.
	finalResult := fmt.Sprintf("%s\n\n--- REVIEW ARQUITETURAL ---\n%s", v2Output, reviewOutput)

	return finalResult, nil
}
