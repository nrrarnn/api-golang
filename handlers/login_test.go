package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"api-golang/models"
	"api-golang/handlers" 
)

func TestLoginHandler(t *testing.T) {
	os.Setenv("JWT_SECRET", "mysecretkey")

	models.Customers = []models.Customer{
		{
			ID:       "1",
			Name:     "testuser",
			Password: "testpass",
		},
	}

	// Skenario 1: Login Berhasil
	t.Run("successful login", func(t *testing.T) {
		loginData := `{"name":"testuser", "password":"testpass"}`
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(loginData)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.LoginHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		expected := "Login successful for customer: testuser"
		if !contains(rr.Body.String(), expected) {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})

	// Skenario 2: Login Gagal Karena Kredensial Tidak Valid
	t.Run("invalid credentials", func(t *testing.T) {
		loginData := `{"name":"wronguser", "password":"wrongpass"}`
		req, err := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(loginData)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.LoginHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusUnauthorized {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
		}
	})

	// Skenario 3: Metode Permintaan Tidak Valid
	t.Run("invalid request method", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/login", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.LoginHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusMethodNotAllowed {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
		}
	})
}

// Fungsi helper untuk memeriksa apakah string mengandung substring
func contains(str, substr string) bool {
	return strings.Contains(str, substr)
}
