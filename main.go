package main

import (
	"net/http"

	"github.com/inconshreveable/log15"
)

func main() {

	log := log15.New()
	http.HandleFunc("/bar", barHandler)

	log.Info("Server Started", "Port", 8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error("Failed to start server", "Error", err)
	}
}
