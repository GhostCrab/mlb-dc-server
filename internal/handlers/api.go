package handlers

import (
	"fmt"
	"net/http"

	"github.com/ghostcrab/mlb-dc-server/internal/tools"
)

func Handler(r *http.ServeMux) {
	// Global middleware
	// r.Use(chimiddle.StripSlashes)

	r.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, tools.DoMongo())
	})
}