package handlers

import (
	"encoding/json"
	"models"
	"net/http"
	"online_shop/database"
)

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	result := database.DB.Find(&products)
	if result.Error != nil {
		http.Error(w, "Could not retrieve products", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)

}
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if result := database.DB.Create(&product); result.Error != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
