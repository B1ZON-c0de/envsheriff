package reporter

import (
	"fmt"
	"io"
	"strings"
)

const (
	ResetColor = "\033[0m"
	RedColor   = "\033[31m"
	GreenColor = "\033[32m"

	SuccessEnvMsg   = "ОК: Все переменные синхронизированы"
	UnsuccessEnvMsg = "Расхождения найдены:\n\nОтсутствуют в .env.example (добавьте их):\n"
)

func PrintCheckedEnv(w io.Writer, analyzedEnv map[string]bool) {
	var unfoundVariables []string

	//Проверяем есть ли ключи со значением false
	for k := range analyzedEnv {
		if !analyzedEnv[k] {
			unfoundVariables = append(unfoundVariables, k)
		}
	}

	//Если всё ок то выводим сообщение об успешной проверке
	if len(unfoundVariables) == 0 {
		fmt.Fprintf(w, "%s%s%s", GreenColor, SuccessEnvMsg, ResetColor)
		return
	}

	//Если были найдены расхождения то выводим их
	unfoundVarString := formatUnfoundVar(unfoundVariables)

	fmt.Fprint(w, unfoundVarString)
}

func formatUnfoundVar(unfoundVariables []string) string {
	var sb strings.Builder

	coloredUnsuccessEnvMsg := fmt.Sprintf("%s%s%s", RedColor, UnsuccessEnvMsg, ResetColor)
	sb.WriteString(coloredUnsuccessEnvMsg)

	for _, k := range unfoundVariables {
		sb.WriteString(fmt.Sprintf("%s- %s\n%s", RedColor, k, ResetColor))
	}

	return sb.String()
}
