package models

import (
	"encoding/json"
	"os"
)

var (
	dataFile = "data/data.json"
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
