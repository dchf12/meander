package main

import (
	"encoding/json"
	"meander/meander"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		_ = respond(w, r, meander.Journeys)
	})
	http.HandleFunc("/recommendations", func(w http.ResponseWriter, r *http.Request) {
		q := &meander.Query{
			Journey: strings.Split(r.URL.Query().Get("journey"), "|"),
		}
		q.Lat, _ = strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		q.Lng, _ = strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
		q.Radius, _ = strconv.Atoi(r.URL.Query().Get("radius"))
		q.CostRangeStr = r.URL.Query().Get("cost")
		places := q.Run()
		_ = respond(w, r, places)
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
