package httpapi

import (
	"flights-master/internal/httpapi/travel"

	"go.uber.org/fx"
)

var Module = fx.Module("httpapi", fx.Provide(
	travel.New,
))
