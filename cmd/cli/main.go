package main

import (
	"envsheriff/internal/analyzer"
	"envsheriff/internal/parser"
	"envsheriff/internal/reporter"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	SuccessSynchronized = "Файлы .env и .env.example успешно синхронизированы"
	HelpMessage         = "Введите:\n -check для проверки .env файлов\n-sync для синхронизации .env и .env.example"
)

func main() {
	args := os.Args[1:]
	HandleCLI(args, os.Stdout)
}

func HandleCLI(args []string, w io.Writer) {
	//Проверка кол-ва аргументов
	if len(args) != 1 {
		fmt.Fprint(w, "Неверное кол-во аргументов")
	}

	//Switch по Имени аргумента
	switch args[0] {
	case "check":
		//Открываем .env файл
		envFile, err := os.Open(".env")
		if err != nil {
			fmt.Errorf("не удалось найти файл .env, %w", err)
			os.Exit(1)
		}
		defer envFile.Close()

		//Открываем .env.example файл
		envExampleFile, err := os.Open(".env.example")
		if err != nil {
			fmt.Errorf("не удалось найти файл .env.example, %w", err)
			os.Exit(1)
		}
		defer envExampleFile.Close()

		//Получаем переменные из .env
		envVariables, err := parser.ParseEnv(envFile)
		if err != nil {
			fmt.Errorf("не удалось получить перемнные из .env: %v", err)
		}

		//Получаем переменные из .env.example
		envExampleVariables, err := parser.ParseEnv(envExampleFile)
		if err != nil {
			fmt.Errorf("не удалось получить перемнные из .env.example: %v", err)

		}

		//Получаем анализ env файлов
		analyzedEnv := analyzer.CompareEnv(envVariables, envExampleVariables)

		//Вывод в консоль
		reporter.PrintCheckedEnv(w, analyzedEnv)
	case "sync":
		//Открываем .env файл
		envFile, err := os.Open(".env")
		if err != nil {
			fmt.Errorf("не удалось найти файл .env, %w", err)
			os.Exit(1)
		}
		defer envFile.Close()

		//Открываем .env.example файл, уже очищенный
		envExampleFile, err := os.OpenFile(".env.example", os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Errorf("не удалось найти файл .env.example, %w", err)
			os.Exit(1)
		}
		defer envExampleFile.Close()

		envVariables, err := parser.ParseEnv(envFile)
		if err != nil {
			fmt.Errorf("не удалось получить перемнные из .env: %v", err)
		}

		for i, envVariable := range envVariables {
			envVariables[i] = envVariable + "="
		}

		var sb strings.Builder
		for _, envVariable := range envVariables {
			sb.WriteString(fmt.Sprintln(envVariable))
		}

		envExampleFile.Write([]byte(sb.String()))

		fmt.Fprint(w, SuccessSynchronized)

	default:
		fmt.Fprint(w, HelpMessage)
	}

}
