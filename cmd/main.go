package main

import (
	"encoding/json"
	"meander/meander"
	"net/http"
)

func main() {
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		_ = respond(w, r, meander.Journeys)
	})

	_ = http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []any) error {
	publicData := make([]any, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}
