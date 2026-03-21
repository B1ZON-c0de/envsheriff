package main

import (
	"bytes"
	"envsheriff/internal/parser"
	"envsheriff/internal/reporter"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestHandlerCLI(t *testing.T) {
	t.Run("Аргумент 'check'(OK)", func(t *testing.T) {
		var buf bytes.Buffer
		mockEnvData := "DEBUG=true\nLOG_LEVEL=INFO"
		clearEnv := createWorkDirEnvFile(t, ".env", mockEnvData)
		defer clearEnv()

		mockEnvExampleData := "DEBUG=true\nLOG_LEVEL=INFO"
		clearEnvExample := createWorkDirEnvFile(t, ".env.example", mockEnvExampleData)
		defer clearEnvExample()

		args := []string{"check"}

		HandleCLI(args, &buf)

		want := fmt.Sprintf("%s%s%s", reporter.GreenColor, reporter.SuccessEnvMsg, reporter.ResetColor)

		if want != buf.String() {
			t.Errorf("Ожидали %q Получили %q", want, buf.String())
		}
	})

	t.Run("Аргумент 'check'(Расхождения)", func(t *testing.T) {
		var buf bytes.Buffer
		mockEnvData := "DEBUG=true\nLOG_LEVEL=INFO"
		clearEnv := createWorkDirEnvFile(t, ".env", mockEnvData)
		defer clearEnv()

		mockEnvExampleData := "DEBUG=true\n"
		clearEnvExample := createWorkDirEnvFile(t, ".env.example", mockEnvExampleData)
		defer clearEnvExample()

		args := []string{"check"}

		HandleCLI(args, &buf)

		want := fmt.Sprintf("%sРасхождения найдены:\n\nОтсутствуют в .env.example (добавьте их):\n%s%s- LOG_LEVEL\n%s", reporter.RedColor, reporter.ResetColor, reporter.RedColor, reporter.ResetColor)

		assertBuf(t, &buf, want)
	})

	t.Run("Аргумент 'sync'", func(t *testing.T) {
		var buf bytes.Buffer

		mockEnvData := "DEBUG=true\nLOG_LEVEL=INFO"
		clearEnv := createWorkDirEnvFile(t, ".env", mockEnvData)
		defer clearEnv()

		mockEnvExampleData := "DEBUG=true"
		clearEnvExample := createWorkDirEnvFile(t, ".env.example", mockEnvExampleData)
		defer clearEnvExample()

		args := []string{"sync"}

		HandleCLI(args, &buf)

		mockEnvExampleFile, err := os.Open(".env.example")
		if err != nil {
			t.Fatal(err)
		}
		defer mockEnvExampleFile.Close()

		got, err := parser.ParseEnv(mockEnvExampleFile)
		if err != nil {
			t.Fatal(err)
		}
		want := []string{"DEBUG", "LOG_LEVEL"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Ожидали %v Получили %v", want, got)
		}

	})
}

func createWorkDirEnvFile(t *testing.T, filename string, initialData string) func() {
	t.Helper()
	homeDir, _ := os.Getwd()
	tempFilePath := filepath.Join(homeDir, filename)

	os.WriteFile(tempFilePath, []byte(initialData), 0644)

	clearFile := func() {
		os.Remove(tempFilePath)
	}
	return clearFile
}

func assertBuf(t *testing.T, buf *bytes.Buffer, want string) {
	t.Helper()

	if want != buf.String() {
		t.Errorf("Ожидали %q Получили %q", want, buf.String())
	}
}
