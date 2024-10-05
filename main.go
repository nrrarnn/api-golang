package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"api-golang/handlers"
	"api-golang/models"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = models.LoadData()
	if err != nil {
		log.Fatalf("Failed to load data: %v", err)
	}

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/payment", handlers.PaymentHandler)

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}