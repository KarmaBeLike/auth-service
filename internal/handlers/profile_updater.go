package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mobidev/internal/models"
	"mobidev/internal/storage"
	"mobidev/internal/utils"

	"github.com/go-playground/validator/v10"
)

func ProfileUpdater(w http.ResponseWriter, r *http.Request) {
	var form models.ProfileUpdater

	// Decoder works only with struct tags
	// and ignores the rest of json fields.
	// TODO: reject wrong structured json string.
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		msg := fmt.Sprintf("Error decoding json: %s", err.Error())

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	validator := validator.New()
	validator.RegisterValidation("date", utils.Bdate)

	if err := validator.Struct(form); err != nil {
		msg := err.Error()

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Check if user email not registred
	user, err := storage.InMemory.Load(form.Email)
	if err == nil {
		msg := "user email already registered"

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	session, err := r.Cookie("session")
	if err != nil {
		msg := "cookies is empty"

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	if valid := storage.InMemory.ValidCookies(session); !valid {
		msg := "cookies expired"

		response := models.Response{Message: msg}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	user.Email = form.Email
	user.Name = form.Name
	user.BirthDate = form.BirthDate
	user.Phone = form.Phone
	
	// !!!FIX: user duplication on save
	storage.InMemory.Save(user)

	msg := "Update successful"
	response := models.Response{Message: msg}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
