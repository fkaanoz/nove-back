package handlers

import (
	"github.com/dimfeld/httptreemux/v5"
	"net/http"
	"shtil/app/shtil/handlers/kafkagrp"
	"shtil/app/shtil/handlers/testgrp"
	"shtil/app/shtil/handlers/usergrp"
	"shtil/business/mids"
	"shtil/business/store/core"
	"shtil/foundation/web"
)

func NewApp(appConfig *web.AppConfig) *web.App {
	app := &web.App{
		ContextMux:  httptreemux.NewContextMux(),
		Logger:      appConfig.Logger,
		ServerErrCh: appConfig.ServerErrCh,
		Middlewares: []web.Middleware{mids.Panic(), mids.Logger(appConfig.Logger), mids.Error()},
		Redis:       appConfig.Redis,
		DB:          appConfig.DB,
	}

	return v1(app)
}

func v1(app *web.App) *web.App {

	userHandlers := usergrp.UserGrp{Redis: app.Redis, Logger: app.Logger, Core: &core.UserCore{DB: app.DB}}

	app.Handle(http.MethodGet, "/:id", userHandlers.UserByID)
	app.Handle(http.MethodPost, "/by-name", userHandlers.UserByName)
	app.Handle(http.MethodGet, "/by-email", userHandlers.UserByEmail)

	// kafka
	app.Handle(http.MethodGet, "/kafka", kafkagrp.Queue)

	// health check - go-api
	app.Handle(http.MethodGet, "/go-api", testgrp.HealthCheck)

	return app
}
