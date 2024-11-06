package main

import (
	"encoding/json"
	"net/http"
)

type ColorResponse struct {
	Color string `json:"color"`
}

func primarioHandler(w http.ResponseWriter, r *http.Request) {
	response := ColorResponse{Color: "blue"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func secundarioHandler(w http.ResponseWriter, r *http.Request) {
	response := ColorResponse{Color: "green"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/primario", primarioHandler)
	http.HandleFunc("/secundario", secundarioHandler)
	http.ListenAndServe(":8080", nil)
}
