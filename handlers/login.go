package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"api-golang/data"
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

	err = data.LoadData()
	if err != nil {
		http.Error(w, "Failed to load data", http.StatusInternalServerError)
		return
	}

	for _, customer := range data.Customers {
		if customer.Name == credentials.Name && customer.Password == credentials.Password {
			fmt.Fprintf(w, "Login successful for customer: %s", customer.Name)

			data.Histories = append(data.Histories, data.History{
				ID:         fmt.Sprintf("%d", len(data.Histories)+1),
				CustomerID: customer.ID,
				Amount:     0,
				Action:     "login",
			})

			if err := data.SaveData(); err != nil {
				http.Error(w, "Failed to save data", http.StatusInternalServerError)
			}

			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
