package handlers

import (
	"fmt"
	"net/http"
	"api-golang/data"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if data.LoggedInCustomer == nil {
		http.Error(w, "No user is logged in", http.StatusUnauthorized)
		return
	}

	data.Histories = append(data.Histories, data.History{
		ID:         fmt.Sprintf("%d", len(data.Histories)+1),
		CustomerID: data.LoggedInCustomer.ID,
		Amount:     0,
		Action:     "logout",
	})

	err := data.SaveData()
	if err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	data.LoggedInCustomer = nil

	fmt.Fprintf(w, "Logout successful")
}
