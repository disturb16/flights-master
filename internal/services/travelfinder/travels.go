package travelfinder

import (
	"context"
	"flights-master/internal/models"
	"flights-master/internal/stores/serpapi"
	"flights-master/logger"
	"sort"
	"sync"
)

type travelManager struct {
	s serpapi.Finder
}

type Manager interface {
	GetBestFlights(ctx context.Context, to, from, date string) ([]models.Flight, error)
	GetBestHotels(ctx context.Context, destination, date string) ([]models.Hotel, error)
	PlanTravel(context.Context, TravelParams) TravelInfo
}

type TravelInfo struct {
	Flights []models.Flight
	Hotels  []models.Hotel
}

type TravelParams struct {
	OriginID      string
	DestinationID string
	Destination   string
	Date          string
}

func New(fs serpapi.Finder) Manager {
	return &travelManager{
		s: fs,
	}
}

func (f *travelManager) GetBestFlights(ctx context.Context, to, from, date string) ([]models.Flight, error) {
	log := logger.FromContext(ctx)

	response, err := f.s.SearchFlights(ctx, to, from, date)
	if err != nil {
		log.WithError(err).Error("failed to get flights")
		return nil, err
	}

	// Sort by price
	sort.Slice(response.BestFlights, func(i, j int) bool {
		prev := response.BestFlights[i]
		next := response.BestFlights[j]
		return prev.Price < next.Price
	})

	flights := []models.Flight{}
	count := min(5, len(response.BestFlights))
	for i := 0; i < count; i++ {
		bestFlight := response.BestFlights[i]
		airline := ""
		if len(bestFlight.Flights) > 0 {
			airline = bestFlight.Flights[0].Airline
		}

		flights = append(flights, models.Flight{
			Airline:  airline,
			Duration: bestFlight.Duration,
			Price:    bestFlight.Price,
		})
	}

	return flights, nil
}

func (tm *travelManager) GetBestHotels(ctx context.Context, destination, date string) ([]models.Hotel, error) {
	log := logger.FromContext(ctx)

	response, err := tm.s.SearchHotels(ctx, destination, date)
	if err != nil {
		log.WithError(err).Error("failed to get flights")
		return nil, nil
	}

	// Sort by price
	sort.Slice(response.BestHotels, func(i, j int) bool {
		prev := response.BestHotels[i]
		next := response.BestHotels[j]
		return prev.Price() < next.Price()
	})

	hotels := []models.Hotel{}

	count := min(5, len(response.BestHotels))
	for i := 0; i < count; i++ {
		bestHotel := response.BestHotels[i]

		hotels = append(hotels, models.Hotel{
			Name:  bestHotel.Name,
			Price: bestHotel.Price(),
			Location: models.Location{
				Latitude:  bestHotel.Location.Latitude,
				Longitude: bestHotel.Location.Longitude,
			},
		})
	}

	return hotels, nil
}

func (tm *travelManager) PlanTravel(ctx context.Context, params TravelParams) TravelInfo {
	log := logger.FromContext(ctx)
	flightsChan := make(chan []models.Flight)
	hotelsChan := make(chan []models.Hotel)
	defer close(flightsChan)
	defer close(hotelsChan)

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		ff, err := tm.GetBestFlights(ctx, params.DestinationID, params.OriginID, params.Date)
		if err != nil {
			log.WithError(err).Error("failed to fetch flights")
		}

		flightsChan <- ff
	}()

	go func() {
		defer wg.Done()
		hh, err := tm.GetBestHotels(ctx, params.Destination, params.Date)
		if err != nil {
			log.WithError(err).Error("failed to fetch hotels")
		}

		hotelsChan <- hh
	}()

	return TravelInfo{
		Flights: <-flightsChan,
		Hotels:  <-hotelsChan,
	}
}
