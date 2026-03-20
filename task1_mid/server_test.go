package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler вернул неверный статус-код: получили %v, ожидали %v",
			status, http.StatusOK)
	}

	expected := "Привет, мир!"
	if rr.Body.String() != expected {
		t.Errorf("handler вернул неожиданное тело ответа: получили %v, ожидали %v",
			rr.Body.String(), expected)
	}
}
