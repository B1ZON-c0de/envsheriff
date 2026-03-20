package reporter_test

import (
	"bytes"
	"envsheriff/internal/reporter"
	"testing"
)

func TestPrintCheckedEnv(t *testing.T) {
	t.Run("Все переменные одинаковы", func(t *testing.T) {
		var buf bytes.Buffer

		mockAnalyzedEnv := map[string]bool{
			"DEBUG": true,
			"INFO":  true,
		}

		reporter.PrintCheckedEnv(&buf, mockAnalyzedEnv)

		want := reporter.SuccessEnvMsg

		assertBuf(t, &buf, want)
	})

	t.Run("Одна перемнная имеет статус false", func(t *testing.T) {
		var buf bytes.Buffer

		mockAnalyzedEnv := map[string]bool{
			"DEBUG": false,
			"INFO":  true,
		}

		reporter.PrintCheckedEnv(&buf, mockAnalyzedEnv)

		want := "Расхождения найдены:\n\nОтсутствуют в .env.example (добавьте их):\n- DEBUG\n"

		assertBuf(t, &buf, want)
	})
}

func assertBuf(t *testing.T, buf *bytes.Buffer, want string) {
	t.Helper()

	if want != buf.String() {
		t.Errorf("Ожидали %q Получили %q", want, buf.String())
	}
}
