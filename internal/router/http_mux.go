package router

import (
	"mobidev/internal/handlers"

	"net/http"
)

// NewRouter creates new request router
// and initialize it with serving paths and
// corresponding handlers.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", handlers.AuthorizationHandler)
	mux.HandleFunc("POST /register", handlers.RegistrationHandler)

	return mux
}
