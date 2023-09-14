package handlers

import (
	"encoding/json"
	"net/http"
	"online_shop/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

type RegistrationInput struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type LoginInput struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	input := RegistrationInput{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     "user",
		IsActive: true,
	}
	if err := DB.Create(&user).Error; err != nil {
		http.Error(w, "Error creating database", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user list"))
	json.NewEncoder(w).Encode(user)

}
