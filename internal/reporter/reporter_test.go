package reporter_test

import (
	"bytes"
	"envsheriff/internal/reporter"
	"fmt"
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

		want := fmt.Sprintf("%s%s%s", reporter.GreenColor, reporter.SuccessEnvMsg, reporter.ResetColor)

		assertBuf(t, &buf, want)
	})

	t.Run("Одна перемнная имеет статус false", func(t *testing.T) {
		var buf bytes.Buffer

		mockAnalyzedEnv := map[string]bool{
			"DEBUG": false,
			"INFO":  true,
		}

		reporter.PrintCheckedEnv(&buf, mockAnalyzedEnv)

		want := fmt.Sprintf("%sРасхождения найдены:\n\nОтсутствуют в .env.example (добавьте их):\n%s%s- DEBUG\n%s", reporter.RedColor, reporter.ResetColor, reporter.RedColor, reporter.ResetColor)

		assertBuf(t, &buf, want)
	})
}

func assertBuf(t *testing.T, buf *bytes.Buffer, want string) {
	t.Helper()

	if want != buf.String() {
		t.Errorf("Ожидали %q Получили %q", want, buf.String())
	}
}
