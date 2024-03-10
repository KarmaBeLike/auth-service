package server

import (
	"mobidev/internal/config"
	"mobidev/internal/router"

	"net/http"
)

// NewServer
func NewServer() *http.Server {
	mux := router.NewRouter()

	server := &http.Server{
		Addr:    config.SEVRER_ADDR,
		Handler: mux,
	}

	return server

}
