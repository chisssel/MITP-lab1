package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainReceiver(t *testing.T) {
	testJson := `{
		"user": {
			"id": 999,
			"username": "tester",
			"email": "test@example.com"
		},
		"payload": "Test Payload",
		"is_valid": false
	}`

	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	oldStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	go func() {
		defer w.Close()
		w.Write([]byte(testJson))
	}()

	main()

	wOut.Close()
	var buf bytes.Buffer
	io.Copy(&buf, rOut)
	capturedOutput := buf.String()

	os.Stdin = oldStdin
	os.Stdout = oldStdout

	expectedUsername := "tester"
	if !strings.Contains(capturedOutput, expectedUsername) {
		t.Errorf("Вывод не содержит ожидаемое имя пользователя: ожидали '%s', получили '%s'", expectedUsername, capturedOutput)
	}

	expectedPayload := "Test Payload"
	if !strings.Contains(capturedOutput, expectedPayload) {
		t.Errorf("Вывод не содержит ожидаемую полезную нагрузку: ожидали '%s', получили '%s'", expectedPayload, capturedOutput)
	}

	expectedConfirmation := "Данные успешно получены"
	if !strings.Contains(capturedOutput, expectedConfirmation) {
		t.Errorf("Вывод не содержит подтверждения об успехе: ожидали '%s', получили '%s'", expectedConfirmation, capturedOutput)
	}
}
