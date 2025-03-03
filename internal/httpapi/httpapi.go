package httpapi

import (
	"flights-master/internal/httpapi/fligts"

	"go.uber.org/fx"
)

var Module = fx.Module("httpapi", fx.Provide(
	fligts.New,
))
