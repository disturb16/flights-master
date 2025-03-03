package fligths

import (
	"context"
	"flights-master/internal/stores/serpapi"
	"flights-master/logger"
)

type fligthsCalculator struct {
	s serpapi.FlightSearcher
}

type Calculator interface {
	GetBestPrice(context.Context)
}

func New(fs serpapi.FlightSearcher) Calculator {
	return &fligthsCalculator{
		s: fs,
	}
}

func (f *fligthsCalculator) GetBestPrice(ctx context.Context) {
	log := logger.FromContext(ctx)

	response, err := f.s.SearchFlights(ctx)
	if err != nil {
		log.WithError(err).Error("failed to get flights")
	}

	log.WithAny("response", response).Info("success")
}
