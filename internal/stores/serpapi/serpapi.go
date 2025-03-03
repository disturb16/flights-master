package serpapi

import (
	"context"
	"flights-master/settings"
)

type serpapi struct {
	apiKey string
}

type Finder interface {
	FindFlights(ctx context.Context, to, from, date string) (SearchResponse, error)
	FindHotels(ctx context.Context, destination, date string) (HotelResponse, error)
}

func New(s *settings.Settings) Finder {
	return &serpapi{
		apiKey: s.Serpapi.ApiKey,
	}
}
