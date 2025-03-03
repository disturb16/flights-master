package stores

import (
	"flights-master/internal/stores/serpapi"

	"go.uber.org/fx"
)

var Module = fx.Module("stores", fx.Provide(
	serpapi.New,
))
