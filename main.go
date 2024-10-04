package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var (
	mu       sync.Mutex
	dataFile = "data.json"
	data     Data
)

type Customer struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Merchant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type History struct {
	ID         string  `json:"id"`
	CustomerID string  `json:"customer_id"`
	Amount     float64 `json:"amount"`
	Action     string  `json:"action"`
}

type Data struct {
	Customers []Customer `json:"customers"`
	Merchants []Merchant `json:"merchants"`
	History   []History  `json:"history"`
}

func loadData() error {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func saveData() error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, file, 0644)
}

var loggedInCustomer *Customer

func loginHandler(w http.ResponseWriter, r *http.Request) {
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

	log.Printf("Loaded Customers: %v", data.Customers)

	for _, customer := range data.Customers {
		if customer.Name == credentials.Name && customer.Password == credentials.Password {
			fmt.Fprintf(w, "Login successful for customer: %s", customer.Name)

			mu.Lock()
			loggedInCustomer = &customer 
			data.History = append(data.History, History{
				ID:         fmt.Sprintf("%d", len(data.History)+1),
				CustomerID: customer.ID,
				Amount:     0,
				Action:     "login",
			})
			err := saveData()
			mu.Unlock()

			if err != nil {
				http.Error(w, "Failed to save data", http.StatusInternalServerError)
				return
			}

			return
		}
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if loggedInCustomer == nil {
		http.Error(w, "No customer is currently logged in", http.StatusBadRequest)
		return
	}

	loggedInCustomer = nil
	fmt.Fprintf(w, "Logout successful")
}

func paymentToMerchantHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if loggedInCustomer == nil {
		http.Error(w, "You must be logged in to make a payment", http.StatusUnauthorized)
		return
	}

	var payment struct {
		MerchantID string  `json:"merchant_id"`
		Amount     float64 `json:"amount"`
	}

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	var merchantFound *Merchant
	for _, merchant := range data.Merchants {
		if merchant.ID == payment.MerchantID {
			merchantFound = &merchant
			break
		}
	}

	if merchantFound == nil {
		http.Error(w, "Merchant not found", http.StatusBadRequest)
		return
	}

	mu.Lock()
	data.History = append(data.History, History{
		ID:         fmt.Sprintf("%d", len(data.History)+1),
		CustomerID: loggedInCustomer.ID,
		Amount:     payment.Amount,
		Action:     "payment to merchant",
	})
	mu.Unlock()

	err = saveData()
	if err != nil {
		http.Error(w, "Failed to save payment history", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Payment to merchant %s successful. Amount: %.2f", merchantFound.Name, payment.Amount)
}


func main() {
	err := loadData()
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/payment", paymentToMerchantHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}