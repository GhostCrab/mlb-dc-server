package handlers

import (
	"net/http"
)

func Handler(r *http.ServeMux) {
	// Global middleware
	// r.Use(chimiddle.StripSlashes)

	r.HandleFunc("/account/{func}", func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("func")

		// Middleware for /account route
		// router.Use(middleware.Authorization)

		w.Write([]byte(id))

		// 'router.Get("/coins", GetCoinBalance)
	})
}