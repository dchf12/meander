package meander

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

//go:embed google_places_key
var APIKey string

type Place struct {
	*GoogleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photos          []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

type googleResponse struct {
	Results []*Place `json:"results"`
}

type GoogleGeometry struct {
	*GoogleLocation `json:"location"`
}

type GoogleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

func (p *Place) Public() any {
	return map[string]any{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photos,
		"vicinity": p.Vicinity,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}

type Query struct {
	Lat          float64
	Lng          float64
	Journey      []string
	Radius       int
	CostRangeStr string
}

func (q *Query) find(types string) (*googleResponse, error) {
	u := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"

	vals := make(url.Values)
	vals.Set("location", fmt.Sprintf("%g,%g", q.Lat, q.Lng))
	vals.Set("radius", fmt.Sprintf("%d", q.Radius))
	vals.Set("type", types)
	vals.Set("key", APIKey)
	if len(q.CostRangeStr) > 0 {
		r := ParseCostRange(q.CostRangeStr)
		vals.Set("minprice", fmt.Sprintf("%d", int(r.From)-1))
		vals.Set("maxprice", fmt.Sprintf("%d", int(r.To)-1))
	}

	res, err := http.Get(u + "?" + vals.Encode())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var r io.Reader = res.Body
	// r = io.TeeReader(r, os.Stderr)
	var response googleResponse
	if err := json.NewDecoder(r).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (q *Query) Run() []any {
	var wg sync.WaitGroup
	places := make([]any, len(q.Journey))
	for i, r := range q.Journey {
		wg.Add(1)
		go func(types string, i int) {
			defer wg.Done()
			resp, err := q.find(types)
			if err != nil {
				return
			}
			if len(resp.Results) == 0 {
				return
			}
			for _, result := range resp.Results {
				for _, photo := range result.Photos {
					photo.URL = strings.TrimRight(fmt.Sprintf("https://maps.googleapis.com/maps/api/place/photo"+
						"?maxwidth=1000&photo_reference=%s&key=%s", photo.PhotoRef, APIKey), "\n")
				}
			}
			places[i] = resp.Results[rand.Intn(len(resp.Results))]
		}(r, i)
	}
	wg.Wait()
	return places
}
