package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"api-golang/models" 
	"github.com/golang-jwt/jwt"
	"os"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var credentials struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = models.LoadData()
	if err != nil {
		http.Error(w, "Failed to load data", http.StatusInternalServerError)
		return
	}

	if len(models.Customers) == 0 {
		http.Error(w, "No customers found", http.StatusInternalServerError)
		return
	}

	for _, customer := range models.Customers {
		if customer.Name == credentials.Name && customer.Password == credentials.Password {

			jwtSecret := []byte(os.Getenv("JWT_SECRET"))
			
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"customer_id": customer.ID,
				"exp":         time.Now().Add(time.Hour * 72).Unix(), // Token expiration time
			})

			tokenString, err := token.SignedString(jwtSecret)
			if err != nil {
				http.Error(w, "Failed to generate token", http.StatusInternalServerError)
				return
			}
			response := struct {
				Message string `json:"message"`
				Token   string `json:"token"`
			}{
				Message: fmt.Sprintf("Login successful for customer: %s", customer.Name),
				Token:   tokenString,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(response)

			models.Histories = append(models.Histories, models.History{
				ID:         fmt.Sprintf("%d", len(models.Histories)+1),
				CustomerID: customer.ID,
				Amount:     0,
				Action:     "login",
			})

			if err := models.SaveData(); err != nil {
				http.Error(w, "Failed to save data", http.StatusInternalServerError)
			}

			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
