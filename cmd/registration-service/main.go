package main

import (
	"mobidev/internal/server"

	"log"
)

func main() {
	srv := server.NewServer()

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("server error: %s\n", err.Error())
	}

}
