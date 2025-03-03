package serpapi

import (
	"context"
	"encoding/json"

	g "github.com/serpapi/google-search-results-golang"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TotalRate struct {
	Lowest int64 `json:"extracted_lowest"`
}

type HotelInfo struct {
	Name      string    `json:"name"`
	Location  Location  `json:"gps_coordinates"`
	TotalRate TotalRate `json:"total_rate"`
}

func (h HotelInfo) Price() int64 {
	return h.TotalRate.Lowest
}

type HotelResponse struct {
	BestHotels []HotelInfo `json:"properties"`
}

func (s *serpapi) SearchHotels(ctx context.Context, destination, date string) (HotelResponse, error) {
	parameter := map[string]string{
		"engine":         "google_hotels",
		"q":              destination,
		"hl":             "en",
		"gl":             "us",
		"check_in_date":  "2025-03-04",
		"check_out_date": "2025-03-05",
		"currency":       "USD",
	}

	response := HotelResponse{}

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
