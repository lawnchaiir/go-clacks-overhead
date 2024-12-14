package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{
			Status: "ok",
		})
	})

	http.ListenAndServe(":8080", &handler{})
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	http.DefaultServeMux.ServeHTTP(w, r)
}
