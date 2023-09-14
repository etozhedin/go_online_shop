package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"online_shop/database"
	"online_shop/models"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all products from the database
	products, err := database.GetAllProducts()

	// Error handling: If there's an error fetching products, log the error and return a 500 status
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Check if the products slice is empty
	if len(products) == 0 {
		w.Write([]byte("No products found"))
		return
	}

	// Send the products as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.AddProduct(&product)
	if err != nil {
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}
	// Here, you'd typically add the product to the database.
	// For now, just return the product in the response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	product, err := database.GetProductById(id)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)

}
func DeleteProductByIdHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID was not found", http.StatusBadRequest)
		return
	}

	err = database.DeleteProductById(id)
	if err != nil {
		http.Error(w, "Product was not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted"))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/ping", PingHandler).Methods("GET")

	r.HandleFunc("/products/{id:[0-9]+}", DeleteProductByIdHandler).Methods("DELETE")
	r.HandleFunc("/products/{id:[0-9]+}", GetProductByIdHandler).Methods("GET")
	r.HandleFunc("/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/products", AddProductHandler).Methods("POST")

	r.HandleFunc("/register", RegisterUser).Methods("POST")
}
