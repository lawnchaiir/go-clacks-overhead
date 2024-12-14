package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
			next.ServeHTTP(w, r)
		})
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		})
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}
