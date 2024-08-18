package main

import "net/http"

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Access-Control-Allow-Origin header to allow all (i.e. no CORS)
		w.Header().Set("Access-Control-Allow-Origin", "http://*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type,X-CSRF-Token, Authorization")
			return
		} else {
			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		}
	})
}
