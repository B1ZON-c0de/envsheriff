package parser

import (
	"os"
	"strings"
)

const (
	separator     = "="
	commentPrefix = "#"
)

func ParseEnv(file *os.File) []string {
	var envVariables []string

	// Получаем значения из .env файла
	data, _ := os.ReadFile(file.Name())
	strData := string(data)

	//Разделение по '\n'
	rawStrSlice := strings.Split(strData, "\n")

	//Отрезаем '=' у каждой переменной
	for _, rawStr := range rawStrSlice {
		if strings.HasPrefix(rawStr, commentPrefix) {
			continue
		}
		envVariables = append(envVariables, sanitizeVariable(rawStr))
	}

	return envVariables
}

// Отрезает у переменной всё после '='
func sanitizeVariable(variable string) string {
	idx := strings.Index(variable, separator)
	return variable[:idx]
}
