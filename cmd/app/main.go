package main

import (
	"context"

	"github.com/Mth-Ryan/waveaction-blog/cmd/app/controllers"
	"github.com/Mth-Ryan/waveaction-blog/cmd/app/views"
	"github.com/Mth-Ryan/waveaction-blog/cmd/app/webserver"
	"github.com/Mth-Ryan/waveaction-blog/internal/application/mappers"
	appservices "github.com/Mth-Ryan/waveaction-blog/internal/application/services"
	cacherepositories "github.com/Mth-Ryan/waveaction-blog/internal/infra/cache-repositories"
	"github.com/Mth-Ryan/waveaction-blog/internal/infra/data"
	eventhandlers "github.com/Mth-Ryan/waveaction-blog/internal/infra/event-handlers"
	"github.com/Mth-Ryan/waveaction-blog/internal/infra/repositories"
	infraservices "github.com/Mth-Ryan/waveaction-blog/internal/infra/services"
	"github.com/Mth-Ryan/waveaction-blog/pkg/conf"
	"github.com/Mth-Ryan/waveaction-blog/pkg/logger"
	"go.uber.org/fx"
)

func RegisterWebServer(lc fx.Lifecycle, ws webserver.WebServer) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go ws.StartServer()
			return nil
		},
		OnStop: func(_ context.Context) error {
			go ws.ShutdownServer()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		conf.Module,
		logger.Module,
		data.Module,
		repositories.Module,
		cacherepositories.Module,
		mappers.Module,
		infraservices.Module,
		eventhandlers.Module,
		appservices.Module,
		views.Module,
		controllers.Module,
		webserver.Module,
		fx.Invoke(RegisterWebServer),
	)
	app.Run()
}
