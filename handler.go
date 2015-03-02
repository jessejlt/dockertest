package main

import (
	"fmt"
	"html"
	"net/http"

	log "github.com/inconshreveable/log15"
)

func barHandler(w http.ResponseWriter, r *http.Request) {

	log.Debug("Request", "Route", "Bar")
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
