package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"api-golang/data"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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

	err = data.LoadData()
	if err != nil {
		http.Error(w, "Failed to load data", http.StatusInternalServerError)
		return
	}

	var validCustomer, validMerchant bool
	for _, customer := range data.Customers {
		if customer.ID == paymentData.CustomerID {
			validCustomer = true
			break
		}
	}
	for _, merchant := range data.Merchants {
		if merchant.ID == paymentData.MerchantID {
			validMerchant = true
			break
		}
	}

	if !validCustomer || !validMerchant {
		http.Error(w, "Invalid customer or merchant", http.StatusBadRequest)
		return
	}

	data.Histories = append(data.Histories, data.History{
		ID:         fmt.Sprintf("%d", len(data.Histories)+1),
		CustomerID: paymentData.CustomerID,
		Amount:     paymentData.Amount,
		Action:     "payment",
	})

	if err := data.SaveData(); err != nil {
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Payment successful for customer: %s", paymentData.CustomerID)
}
