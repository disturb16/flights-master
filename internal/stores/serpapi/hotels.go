package serpapi

import (
	"context"
	"encoding/json"
	"time"

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

const reservationDaysAmount = 5

func (h HotelInfo) Price() int64 {
	return h.TotalRate.Lowest
}

type HotelResponse struct {
	BestHotels []HotelInfo `json:"properties"`
}

func (s *serpapi) FindHotels(ctx context.Context, destination, date string) (HotelResponse, error) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return HotelResponse{}, err
	}

	parameter := map[string]string{
		"engine":         "google_hotels",
		"q":              destination,
		"hl":             "en",
		"gl":             "us",
		"check_in_date":  date,
		"check_out_date": d.AddDate(0, 0, reservationDaysAmount).Format("2006-01-02"),
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
