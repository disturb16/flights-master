package main

import (
	"context"
	"flights-master/database"
	"flights-master/internal/httpapi"
	"flights-master/internal/httpapi/handlers"
	"flights-master/internal/services"
	"flights-master/internal/stores"
	"flights-master/logger"
	"flights-master/settings"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Lc       fx.Lifecycle
	Settings *settings.Settings
	DB       *sqlx.DB
	Router   *echo.Echo
	Logger   *logger.Logger
	Handlers []handlers.Handler `group:"handlers"`
}

func setLifeCycle(p Params) {
	p.Lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			for _, h := range p.Handlers {
				h.RegisterRoutes(p.Router)
			}

			go func() {
				err := database.PopulateDb(p.DB)
				if err != nil {
					p.Logger.WithError(err).Error("failed to populate database")
				}
			}()

			go func() {
				err := http.ListenAndServe(p.Settings.AppAddr, p.Router)
				if err != nil {
					p.Logger.WithError(err).Error(err.Error())
				}
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {
			if err := p.DB.Close(); err != nil {
				p.Logger.WithError(err).Error("failed to close db")
			}

			return nil
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			settings.New,
			logger.New,
			echo.New,
			database.New,
		),
		stores.Module,
		services.Module,
		httpapi.Module,
		fx.Invoke(
			setLifeCycle,
		),
	)

	app.Run()
}
