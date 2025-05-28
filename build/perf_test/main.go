package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	usersCount    = 1000
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
	file, err := os.Create(fmt.Sprintf("%s/%s", targetsDir, signupTargets))
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 1; i <= usersCount; i++ {
		data := map[string]string{
			"login":        fmt.Sprintf("testuser%d", i),
			"password":     "TestPassword123!",
			"phone_number": fmt.Sprintf("+7%010d", i),
			"first_name":   fmt.Sprintf("User%d", i),
			"last_name":    fmt.Sprintf("Lastname%d", i),
		}
		jsonData, _ := json.Marshal(data)

		// ВНИМАНИЕ: нет пустой строки между заголовками и телом
		target := fmt.Sprintf("POST https://%s/api/auth/signup\n", apiAddress) +
			"Content-Type: application/json\n" +
			string(jsonData) + "\n"

		if _, err := file.WriteString(target); err != nil {
			return err
		}
	}
	fmt.Printf("Generated %d signup targets in %s\n", usersCount, signupTargets)
	return nil
}



func generateAuthTargets() error {
	file, err := os.Create(fmt.Sprintf("%s/%s", targetsDir, authTargets))
	if err != nil {
		return err
	}
	defer file.Close()

	for i := 1; i <= usersCount; i++ {
		data := map[string]string{
			"login":    fmt.Sprintf("testuser%d", i),
			"password": "TestPassword123!",
		}
		jsonData, _ := json.Marshal(data)

		target := fmt.Sprintf("POST https://%s/api/auth/signin\n", apiAddress) +
			"Content-Type: application/json\n" +
			"\n" +
			string(jsonData) + "\n\n"

		if _, err := file.WriteString(target); err != nil {
			return err
		}
	}
	fmt.Printf("Generated %d auth targets in %s\n", usersCount, authTargets)
	return nil
}