package main

import (
	"bytes"
	"envsheriff/internal/reporter"
	"fmt"
	"os"
	"path/filepath"
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

		if want != buf.String() {
			t.Errorf("Ожидали %q Получили %q", want, buf.String())
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
