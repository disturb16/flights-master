package serpapi

import (
	"context"
	"flights-master/settings"
)

type serpapi struct {
	apiKey string
}

type FlightSearcher interface {
	SearchFlights(ctx context.Context) (SearchResponse, error)
}

func New(s *settings.Settings) FlightSearcher {
	return &serpapi{
		apiKey: s.Serpapi.ApiKey,
	}
}
