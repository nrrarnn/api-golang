package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"api-golang/models"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Mengambil token JWT dari header Authorization
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Authorization token is missing", http.StatusUnauthorized)
		return
	}

	// Memverifikasi token JWT di sini (tambahkan logika verifikasi token Anda)

	var paymentData struct {
		CustomerID string  `json:"customer_id"`
		Amount     float64 `json:"amount"`
		MerchantID string  `json:"merchant_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&paymentData)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = models.LoadData()
	if err != nil {
		http.Error(w, "Failed to load data", http.StatusInternalServerError)
		return
	}

	var validCustomer, validMerchant bool
	for _, customer := range models.Customers {
		if customer.ID == paymentData.CustomerID {
			validCustomer = true
			break
		}
	}
	for _, merchant := range models.Merchants {
		if merchant.ID == paymentData.MerchantID {
			validMerchant = true
			break
		}
	}

	if !validCustomer || !validMerchant {
		http.Error(w, "Invalid customer or merchant", http.StatusBadRequest)
		return
	}

	models.Histories = append(models.Histories, models.History{
		ID:         fmt.Sprintf("%d", len(models.Histories)+1),
		CustomerID: paymentData.CustomerID,
		Amount:     paymentData.Amount,
		Action:     "payment",
	})

	if err := models.SaveData(); err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Payment successful for customer: %s", paymentData.CustomerID)
}
