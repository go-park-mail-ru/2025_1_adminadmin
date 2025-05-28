package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	usersCount    = 6000
	targetsDir    = "docs/perf_test"
	signupTargets = "signup-targets.txt"
	authTargets   = "auth-targets.txt"
	apiAddress    = "doordashers.ru"
)

func main() {
	if err := os.MkdirAll(targetsDir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	if err := generateSignupTargets(); err != nil {
		fmt.Printf("Error generating signup targets: %v\n", err)
	}

	if err := generateAuthTargets(); err != nil {
		fmt.Printf("Error generating auth targets: %v\n", err)
	}
}

func generateSignupTargets() error {
	bodiesDir := fmt.Sprintf("%s/bodies", targetsDir)
	if err := os.MkdirAll(bodiesDir, os.ModePerm); err != nil {
		return fmt.Errorf("не удалось создать директорию с JSON: %w", err)
	}

	targetsPath := fmt.Sprintf("%s/%s", targetsDir, signupTargets)
	targetsFile, err := os.Create(targetsPath)
	if err != nil {
		return fmt.Errorf("не удалось создать файл целей: %w", err)
	}
	defer targetsFile.Close()

	// Русские имена и фамилии для валидации
	russianFirstNames := []string{"Алексей", "Ирина", "Дмитрий", "Екатерина", "Никита", "Ольга", "Сергей", "Анна", "Иван", "Мария"}
	russianLastNames := []string{"Иванов", "Петрова", "Сидоров", "Кузнецова", "Смирнов", "Попова", "Козлов", "Морозова", "Новиков", "Федорова"}

	for i := 1; i <= usersCount; i++ {
		body := map[string]string{
			"login":        fmt.Sprintf("testuser%d", i),
			"password":     "TestPassword123!",
			"phone_number": fmt.Sprintf("7%010d", i),
			// Используем русские имена и фамилии по циклу
			"first_name": russianFirstNames[(i-1)%len(russianFirstNames)],
			"last_name":  russianLastNames[(i-1)%len(russianLastNames)],
		}

		jsonBody, err := json.MarshalIndent(body, "", "  ")
		if err != nil {
			fmt.Printf("Ошибка сериализации JSON: %v\n", err)
			continue
		}

		jsonFilename := fmt.Sprintf("signup_user_%05d.json", i)
		jsonPath := fmt.Sprintf("%s/%s", bodiesDir, jsonFilename)
		if err := os.WriteFile(jsonPath, jsonBody, 0644); err != nil {
			fmt.Printf("Ошибка записи JSON: %v\n", err)
			continue
		}

		target := fmt.Sprintf("POST https://%s/api/auth/signup\n@%s\n", apiAddress, jsonPath)
		if _, err := targetsFile.WriteString(target); err != nil {
			fmt.Printf("Ошибка записи в targets файл: %v\n", err)
		}
	}

	fmt.Printf("✔️  Сгенерировано %d signup-запросов в %s и тела в %s\n", usersCount, signupTargets, bodiesDir)
	return nil
}

func generateAuthTargets() error {
	bodiesDir := fmt.Sprintf("%s/bodies", targetsDir)
	if err := os.MkdirAll(bodiesDir, os.ModePerm); err != nil {
		return fmt.Errorf("не удалось создать директорию с JSON: %w", err)
	}

	targetsPath := fmt.Sprintf("%s/%s", targetsDir, authTargets)
	targetsFile, err := os.Create(targetsPath)
	if err != nil {
		return fmt.Errorf("не удалось создать файл целей: %w", err)
	}
	defer targetsFile.Close()

	for i := 1; i <= usersCount; i++ {
		body := map[string]string{
			"login":    fmt.Sprintf("testuser%d", i),
			"password": "TestPassword123!",
		}

		jsonBody, err := json.MarshalIndent(body, "", "  ")
		if err != nil {
			fmt.Printf("Ошибка сериализации JSON: %v\n", err)
			continue
		}

		jsonFilename := fmt.Sprintf("auth_user_%05d.json", i)
		jsonPath := fmt.Sprintf("%s/%s", bodiesDir, jsonFilename)
		if err := os.WriteFile(jsonPath, jsonBody, 0644); err != nil {
			fmt.Printf("Ошибка записи JSON: %v\n", err)
			continue
		}

		target := fmt.Sprintf("POST https://%s/api/auth/signin\n@%s\n", apiAddress, jsonPath)
		if _, err := targetsFile.WriteString(target); err != nil {
			fmt.Printf("Ошибка записи в targets файл: %v\n", err)
		}
	}

	fmt.Printf("✔️  Сгенерировано %d auth-запросов в %s и тела в %s\n", usersCount, authTargets, bodiesDir)
	return nil
}
