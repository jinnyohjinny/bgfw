package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Tools struct {
	Name string `json:"name"`
	Cmd  string `json:"cmd"`
}

func ParseJson(filePath string, input string, mode string) (string, error) {
	var tool Tools

	outputDomainPath := fmt.Sprintf("output/%s/%s", input, mode)
	if err := os.MkdirAll(outputDomainPath, os.ModePerm); err != nil {
		return "", fmt.Errorf("gagal membuat direktori: %s, error: %v", outputDomainPath, err)
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("gagal membaca file: %s, error: %v", filePath, err)
	}

	if err := json.Unmarshal(fileContent, &tool); err != nil {
		return "", fmt.Errorf("gagal decode JSON di file: %s, error: %v", filePath, err)
	}

	tool.Cmd = strings.Replace(tool.Cmd, "input", input, -1)
	tool.Cmd = strings.Replace(tool.Cmd, "target", input, -1)

	return tool.Cmd, nil
}
