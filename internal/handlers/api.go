package handlers

import (
	"fmt"
	"net/http"
)

func Handler(r *http.ServeMux) {
	// Global middleware
	// r.Use(chimiddle.StripSlashes)

	r.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the HomePage!")
	})
}