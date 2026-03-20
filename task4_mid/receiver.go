package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type ReceivedData struct {
	User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"user"`
	Payload string `json:"payload"`
	IsValid bool   `json:"is_valid"`
}

func main() {
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Ошибка чтения из stdin: %v", err)
	}

	var data ReceivedData

	if err := json.Unmarshal(inputBytes, &data); err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v", err)
	}

	fmt.Printf("✅ Данные успешно получены и разобраны в Go!\n")
	fmt.Printf("   - Имя пользователя: %s (ID: %d)\n", data.User.Username, data.User.ID)
	fmt.Printf("   - Полезная нагрузка: \"%s\"\n", data.Payload)
	fmt.Printf("   - Валидность: %t\n", data.IsValid)
}
