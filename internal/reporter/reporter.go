package reporter

import (
	"fmt"
	"io"
	"strings"
)

const (
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
		fmt.Fprint(w, SuccessEnvMsg)
		return
	}

	//Если были найдены расхождения то выводим их
	unfoundVarString := formatUnfoundVar(unfoundVariables)

	fmt.Fprint(w, unfoundVarString)
}

func formatUnfoundVar(unfoundVariables []string) string {
	var sb strings.Builder
	sb.WriteString(UnsuccessEnvMsg)

	for _, k := range unfoundVariables {
		sb.WriteString(fmt.Sprintf("- %s\n", k))
	}

	return sb.String()
}
