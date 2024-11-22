package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ColorResponse struct {
	Color string `json:"color"`
}

func primarioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Obteniendo color primario: blue")
	response := ColorResponse{Color: "blue"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func secundarioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Obteniendo color secundario: green")
	response := ColorResponse{Color: "green"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/primario", primarioHandler)
	http.HandleFunc("/secundario", secundarioHandler)
	fmt.Println("Starting server on :8081")
	http.ListenAndServe(":8081", nil)
}
