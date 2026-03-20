package analyzer

func CompareEnv(env, envExample []string) map[string]bool {
	analyzedEnv := make(map[string]bool)

	// Заполняем map переменными из .env.example
	for _, variable := range envExample {
		analyzedEnv[variable] = true
	}

	// Проверяем есть ли переменные в .env
	for _, variable := range env {
		if _, ok := analyzedEnv[variable]; !ok {
			analyzedEnv[variable] = false
		}
	}

	return analyzedEnv
}
