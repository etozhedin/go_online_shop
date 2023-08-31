package main

import (
	"log"
	"net/http"

	"online_shop/api/handlers"
	"online_shop/database"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
func main() {
	database.Connect()
	defer database.Close()

	products, err := database.GetAllProducts()
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}

	for _, product := range products {
		log.Printf("ID: %d, Name: %s, Description: %s, Price: %f, Image: %s", product.ID, product.Name, product.Description, product.Price, product.ImageURL)
	}

	r := mux.NewRouter()
	handlers.RegisterRoutes(r)
	port := "8080"
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
