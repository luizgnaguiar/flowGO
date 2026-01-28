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

	// 1. Inicializa o provedor externo
	gemini, err := provider.NewGeminiClient(ctx, os.Getenv("GEMINI_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	// 2. Inicializa a engine com a lógica de prompt
	engine := &processor.FlowEngine{Provider: gemini}

	// 3. Simula um user input
	userInput := "Como estruturar um projeto Go com 5 camadas?"

	fmt.Println("🤖 FlowGO a processar...")
	result, err := engine.Process(ctx, userInput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n--- RESULTADO FINAL ---\n%s\n", result)
}
