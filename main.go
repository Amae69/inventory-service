package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode([]string{
			"Laptop", "Keyboard", "Phone",
		})
	})
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
