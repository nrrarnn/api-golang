package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"api-golang/models"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "No token provided", http.StatusUnauthorized)
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	models.BlacklistToken(tokenString)

	fmt.Fprintf(w, "Logout successful")
}
