package handlers

import (
	"mobidev/internal/models"
	"mobidev/internal/storage"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const HASHING_COST = 12

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	var form models.RegistrationForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		msg := fmt.Sprintf("Error decoding json: %s", err.Error())

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := validator.New().Struct(form); err != nil {
		msg := err.Error()

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(response)
		return
	}

	if user, exists := storage.InMemory[form.Email]; exists {
		msg := fmt.Sprintf("User %s already exists. Use different email address", user.Email)

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(409)
		json.NewEncoder(w).Encode(response)
		return
	}

	passwordHah, err := bcrypt.GenerateFromPassword([]byte(form.Password), HASHING_COST)
	if err != nil {
		msg := "Can't register right now. Try again after a while"

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Email:         form.Email,
		PasswordHash:  string(passwordHah),
		Name:          "",
		Phone:         "",
		BirthDate:     "",
		Cookie:        uuid.NewString(),
		CookieExpires: time.Now().Add(3 * 24 * time.Hour).Unix(),
	}

	storage.InMemory[form.Email] = user

	msg := "Registration successful!"
	response := models.Response{Message: msg}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
