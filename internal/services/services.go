package services

import (
	"flights-master/internal/services/travelfinder"

	"go.uber.org/fx"
)

var Module = fx.Module("services", fx.Provide(
	travelfinder.New,
))
