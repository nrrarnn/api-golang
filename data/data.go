package data

import (
	"encoding/json"
	"os"
)

var (
	dataFile          = "data.json"
	LoggedInCustomer   *Customer
	Customers          []Customer
	Merchants          []Merchant
	Histories          []History
)

type Data struct {
	Customers []Customer `json:"customers"`
	Merchants []Merchant `json:"merchants"`
	Histories []History  `json:"histories"`
}

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

func LoadData() error {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return err
	}

	var data struct {
		Customers []Customer `json:"customers"`
		Merchants []Merchant `json:"merchants"`
		Histories []History  `json:"histories"`
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return err
	}

	Customers = data.Customers
	Merchants = data.Merchants
	Histories = data.Histories

	return nil
}

func SaveData() error {
	data := struct {
		Customers []Customer `json:"customers"`
		Merchants []Merchant `json:"merchants"`
		Histories []History  `json:"histories"`
	}{
		Customers: Customers,
		Merchants: Merchants,
		Histories: Histories,
	}

	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, file, 0644)
}
