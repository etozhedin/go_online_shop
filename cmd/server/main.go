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
	defer database.GetDB().Close()
	r := mux.NewRouter()
	handlers.RegisterRoutes(r)
	port := "8080"
	log.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
