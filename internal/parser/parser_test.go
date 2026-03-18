package parser_test

import (
	"envsheriff/internal/parser"
	"os"
	"reflect"
	"testing"
)

func TestParseEnv(t *testing.T) {
	t.Run("Получение переменных из .env файла", func(t *testing.T) {
		mockEnvData := "#комментарий\nDEBUG=true\nLOG_LEVEL=info"

		tempEnvFile, clearTempEnvFile := createTempEnvFile(t, ".env", mockEnvData)
		defer clearTempEnvFile()

		got := parser.ParseEnv(tempEnvFile)

		want := []string{
			"DEBUG",
			"LOG_LEVEL",
		}

		assertSlices(t, want, got)
	})
}

func createTempEnvFile(t *testing.T, filename string, initialData string) (*os.File, func()) {
	t.Helper()
	tempFile, _ := os.CreateTemp("", filename)

	tempFile.Write([]byte(initialData))

	clearFile := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}
	return tempFile, clearFile
}

func assertSlices(t *testing.T, want, got []string) {
	t.Helper()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Ожидали %v, Получили %v", want, got)
	}
}
