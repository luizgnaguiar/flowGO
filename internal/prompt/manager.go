package prompt

import (
	"bytes"
	"os"
	"text/template"
)

func BuildPrompt(templateName string, data interface{}) (string, error) {
	// Lê o arquivo .tmpl
	content, err := os.ReadFile("templates/" + templateName)
	if err != nil {
		return "", err
	}

	// Faz o parse do template
	tmpl, err := template.New("prompt").Parse(string(content))
	if err != nil {
		return "", err
	}

	// Executa o template injetando os dados
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
