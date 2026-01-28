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

	// 1. Setup do Provedor
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("ERRO: GEMINI_API_KEY não configurada.")
	}

	gemini, err := provider.NewGeminiClient(ctx, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Inicializar a Engine
	engine := &processor.FlowEngine{Provider: gemini}

	// 3. Input do Utilizador
	userInput := "Sistema de pagamentos em Go com processamento assíncrono."

	fmt.Println("🌊 flowGO Pipeline: Iniciando 3 Camadas de Inteligência")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("📥 ENTRADA: %s\n", userInput)
	fmt.Println("--------------------------------------------------")

	// 4. Execução do Fluxo
	// A engine agora orquestra internamente: Author -> Simulator -> Enforcer
	fmt.Println("🧠 Camada 1: Traduzindo intenção (Prompt Author)...")
	fmt.Println("🔍 Camada 2: Simulando riscos de execução (SRE Simulator)...")
	fmt.Println("🛡️  Camada 3: Aplicando restrições técnicas (Constraint Enforcer)...")

	finalOutput, err := engine.Process(ctx, userInput)
	if err != nil {
		log.Fatalf("Erro no pipeline: %v", err)
	}

	// 5. Resultado Final
	fmt.Println("\n✅ CONTRATO TÉCNICO FINALIZADO (Prompt_v2):")
	fmt.Println("==================================================")
	fmt.Println(finalOutput)
	fmt.Println("==================================================")
}
