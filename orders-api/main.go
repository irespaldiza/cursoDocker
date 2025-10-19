package main

import (
	"encoding/json"
	"net/http"
)

type Order struct {
	ID   int    `json:"id"`
	Item string `json:"item"`
	User string `json:"user"`
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		orders := []Order{
			{1, "Café", "Alice"},
			{2, "Té", "Bob"},
		}
		json.NewEncoder(w).Encode(orders)
	})

	http.ListenAndServe(":8082", nil)
}
