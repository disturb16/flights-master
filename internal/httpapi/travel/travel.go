package travel

import (
	"flights-master/internal/httpapi/handlers"
	"flights-master/internal/services/travelfinder"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type handler struct {
	travelManager travelfinder.Manager
}

type Result struct {
	fx.Out

	Handler handlers.Handler `group:"handlers"`
}

func New(fc travelfinder.Manager) Result {
	return Result{
		Handler: &handler{
			travelManager: fc,
		},
	}
}

func (h *handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/travel", h.SearchBestPrices)
}
