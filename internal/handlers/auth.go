package handlers

import (
	"encoding/json"
	"fmt"
	"mobidev/internal/models"
	"mobidev/internal/storage"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	var form models.AuthenticationForm

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		msg := fmt.Sprintf("Error decoding json: %s", err.Error())

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if client send wrong data
	if err := validator.New().Struct(form); err != nil {
		msg := err.Error()

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if user email not registred
	user, err := storage.InMemory.Load(form.Email)
	if err != nil {
		msg := err.Error()

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Chech authorization
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(form.Password)); err != nil {
		msg := "Wrong password"

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	sessionVal, expires := uuid.NewString(), 3*24*60*60

	user.Cookie = sessionVal
	user.CookieExpires = int64(expires)

	storage.InMemory.Save(user)

	HTTPcookie := http.Cookie{
		Name:     "session",
		Value:    sessionVal,
		Path:     "/",
		MaxAge:   expires,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &HTTPcookie)

	msg := "Authorized successful"
	response := models.Response{Message: msg}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
