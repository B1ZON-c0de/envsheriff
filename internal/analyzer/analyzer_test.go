package analyzer_test

import (
	"envsheriff/internal/analyzer"
	"reflect"
	"testing"
)

func TestCompareEnv(t *testing.T) {
	t.Run("Получить отсутсвующую переменную в .env.example из .env", func(t *testing.T) {
		mockEnv := []string{
			"DEBUG",
			"LOG_LEVEL",
		}

		mockEnvExample := []string{
			"LOG_LEVEL",
		}

		got := analyzer.CompareEnv(mockEnv, mockEnvExample)

		want := map[string]bool{
			"DEBUG":     false,
			"LOG_LEVEL": true,
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Ожидали %v Получили %v", want, got)
		}
	})
}
