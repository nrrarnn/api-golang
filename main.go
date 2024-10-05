package main

import (
	"fmt"
	"log"
	"net/http"
	"api-golang/handlers"
	"api-golang/data"
)

func main() {
	err := data.LoadData()
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/payment", handlers.PaymentHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}