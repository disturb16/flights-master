package serpapi

import (
	"context"
	"encoding/json"

	g "github.com/serpapi/google-search-results-golang"
)

type Airport struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Time string `json:"time"`
}

type FlightStop struct {
	DeparturAirport Airport `json:"departure_airport"`
	ArrivalAirport  Airport `json:"arrival_airport"`
	Airplane        string  `json:"airplane"`
	Airline         string  `json:"airline"`
	TravelClass     string  `json:"travel_class"`
	Number          string  `json:"flight_number"`
	Duration        int64   `json:"duration"`
}

type FlightInfo struct {
	Type       string       `json:"type"`
	Extensions []string     `json:"extensions"` // e.g. legroom, wifi, carbon emssions
	Flights    []FlightStop `json:"flights"`
	Duration   int64        `json:"total_duration"`
	Price      int64        `json:"price"`
}

type SearchResponse struct {
	BestFlights []FlightInfo `json:"best_flights"`
}

func (s *serpapi) SearchFlights(ctx context.Context) (SearchResponse, error) {
	response := SearchResponse{}
	parameter := map[string]string{
		"engine":        "google_flights",
		"departure_id":  "CDG",
		"arrival_id":    "AUS",
		"hl":            "en",
		"gl":            "us",
		"currency":      "USD",
		"outbound_date": "2025-03-03",
		"return_date":   "2025-03-09",
	}

	search := g.NewGoogleSearch(parameter, s.apiKey)
	results, err := search.GetJSON()
	if err != nil {
		return response, err
	}

	bb, err := json.Marshal(results)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(bb, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
