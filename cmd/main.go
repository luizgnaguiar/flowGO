// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"flowGO/internal/processor"
// 	"flowGO/internal/provider"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	// 1. Setup inicial
// 	godotenv.Load()
// 	ctx := context.Background()

// 	apiKey := os.Getenv("GEMINI_API_KEY")
// 	if apiKey == "" {
// 		log.Fatal("ERRO: GEMINI_API_KEY não encontrada no ambiente ou .env")
// 	}

// 	// 2. Inicialização do Cliente Gemini
// 	gemini, err := provider.NewGeminiClient(ctx, apiKey)
// 	if err != nil {
// 		log.Fatalf("Falha ao conectar com Gemini: %v", err)
// 	}

// 	// 3. Inicialização da Engine
// 	engine := &processor.FlowEngine{Provider: gemini}

// 	// 4. Input de teste (Desafio arquitetural)
// 	userInput := "Sistema distribuído de reserva de assentos em tempo real para estádios."

// 	fmt.Println("🌊 flowGO Pipeline: Executando Fluxo de 4 Especialistas")
// 	fmt.Println("---------------------------------------------------------")
// 	fmt.Printf("📥 INPUT DO USUÁRIO: %s\n", userInput)
// 	fmt.Println("---------------------------------------------------------")

// 	// 5. Orquestração das Camadas
// 	fmt.Println("✍️  [1/4] Authoring: Criando Prompt_v1...")
// 	fmt.Println("🕵️  [2/4] Simulator: Identificando Riscos de Runtime...")
// 	fmt.Println("🛡️  [3/4] Enforcer:  Gerando Contrato Técnico (Prompt_v2)...")
// 	fmt.Println("⚖️  [4/4] Reviewer:  Auditando Arquitetura Final...")

// 	finalOutput, err := engine.Process(ctx, userInput)
// 	if err != nil {
// 		log.Fatalf("Erro no Pipeline: %v", err)
// 	}

// 	// 6. Output Final consolidado
// 	fmt.Println("\n🚀 ENTREGA FINAL DO FLOW:")
// 	fmt.Println("=========================================================")
// 	fmt.Println(finalOutput)
// 	fmt.Println("=========================================================")
// }

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"flowGO/internal/processor"
	"flowGO/internal/provider"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	ctx := context.Background()

	gemini, err := provider.NewGeminiClient(ctx, os.Getenv("GEMINI_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	engine := &processor.FlowEngine{Provider: gemini}

	userInput := "Microserviço de autenticação usando JWT e integração com OAuth2 em Go."

	fmt.Println("🚀 flowGO: Executando Pipeline Completo de 5 Camadas")
	fmt.Println("---------------------------------------------------------")

	// Feedback visual do progresso
	stages := []string{
		"✍️  [1/5] Authoring",
		"🕵️  [2/5] Risk Simulation",
		"🛡️  [3/5] Constraint Enforcement",
		"⚖️  [4/5] Architecture Review",
		"🧪 [5/5] Test Designing",
	}

	for _, stage := range stages {
		fmt.Println(stage + "...")
	}

	finalOutput, err := engine.Process(ctx, userInput)
	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	fmt.Println("\n📦 ENTREGA FINAL (Blueprint Completo):")
	fmt.Println("=========================================================")
	fmt.Println(finalOutput)
	fmt.Println("=========================================================")
}
