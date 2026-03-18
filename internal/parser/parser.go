package parser

import (
	"errors"
	"os"
	"strings"
)

const (
	separator     = "="
	commentPrefix = "#"
	space         = " "
	tab           = "\t"

	ErrReadFile = "не удалось прочитать файл"
)

func ParseEnv(file *os.File) ([]string, error) {
	var envVariables []string

	// Получаем значения из .env файла
	data, err := os.ReadFile(file.Name())
	if err != nil {
		return envVariables, errors.New(ErrReadFile)
	}
	strData := string(data)

	//Разделение по '\n'
	rawStrSlice := strings.Split(strData, "\n")

	//Отрезаем '=' у каждой переменной порпуская ' ', '\t'  комментарии
	for _, rawStr := range rawStrSlice {
		switch rawStr {
		case space, tab:
			continue
		default:
			if strings.HasPrefix(rawStr, commentPrefix) {
				continue
			}
			envVariables = append(envVariables, sanitizeVariable(rawStr))
		}

	}

	return envVariables, nil
}

// Отрезает у переменной всё после '='
func sanitizeVariable(variable string) string {
	idx := strings.Index(variable, separator)
	return variable[:idx]
}
