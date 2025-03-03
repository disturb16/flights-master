package services

import (
	"flights-master/internal/services/fligths"

	"go.uber.org/fx"
)

var Module = fx.Module("services", fx.Provide(
	fligths.New,
))
