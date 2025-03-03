package fligts

import (
	"flights-master/internal/httpapi/handlers"
	"flights-master/internal/services/fligths"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	fligthsCalculator fligths.Calculator
}

type Result struct {
	fx.Out

	Handler handlers.Handler `group:"handlers"`
}

func New(fc fligths.Calculator) Result {
	return Result{
		Handler: &handler{
			fligthsCalculator: fc,
		},
	}
}

func (h *handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/flights", h.SearchBestPrices)
}

func (h *handler) SearchBestPrices(c echo.Context) error {
	h.fligthsCalculator.GetBestPrice(c.Request().Context())
	return c.JSON(http.StatusOK, "")
}
