package travel

import (
	"flights-master/internal/services/travelfinder"
	"flights-master/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Params struct {
	OriginAirportCode      string `query:"origin_airport_code"`
	DestinationAirportCode string `query:"destination_airport_code"`
	Destination            string `query:"Destination"`
	Date                   string `query:"date"`
}

func (h *handler) SearchBestPrices(c echo.Context) error {
	ctx := c.Request().Context()
	params := Params{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid parameters")
	}

	log := logger.FromContext(ctx)
	log.WithAny("params", params).Info("")

	info := h.travelManager.PlanTravel(ctx, travelfinder.TravelParams{
		OriginID:      params.OriginAirportCode,
		DestinationID: params.DestinationAirportCode,
		Destination:   params.Destination,
		Date:          params.Date,
	})

	return c.JSON(http.StatusOK, info)
}
