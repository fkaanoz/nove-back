package handlers

import (
	"github.com/dimfeld/httptreemux/v5"
	"net/http"
	"shtil/app/shtil/handlers/logingrp"
	"shtil/app/shtil/handlers/ordergrp"
	"shtil/app/shtil/handlers/system"
	"shtil/app/shtil/handlers/transactiongrp"
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
		Auth:        appConfig.Auth,
	}

	return v1(app)
}

func v1(app *web.App) *web.App {

	// health check for load balancer or k8s
	app.Handle(http.MethodGet, "/api/health-check", system.HealthCheck)

	// user handlers
	userHandlers := usergrp.UserGrp{Redis: app.Redis, Logger: app.Logger, Core: &core.UserCore{DB: app.DB}, Auth: app.Auth}
	app.Handle(http.MethodGet, "/api/user-by-id/:id", userHandlers.UserByID, mids.Auth(app.Auth))
	app.Handle(http.MethodPost, "/api/by-name", userHandlers.UserByName)
	app.Handle(http.MethodGet, "/api/by-email", userHandlers.UserByEmail)

	// order handlers
	orderHandlers := ordergrp.OrderGrp{Logger: app.Logger, Core: &core.OrderCore{DB: app.DB}}
	app.Handle(http.MethodGet, "/api/last-20-orders", orderHandlers.Last20Orders)

	// transaction handlers
	transactionHandlers := transactiongrp.TransactionGrp{Core: &core.TransactionCore{DB: app.DB}}
	app.Handle(http.MethodGet, "/api/tx/:id", transactionHandlers.GetByID)

	// login handlers
	loginHandlers := logingrp.LoginGrp{Core: core.LoginCore{DB: app.DB}, Auth: app.Auth}
	app.Handle(http.MethodPost, "/api/register", loginHandlers.Register)
	app.Handle(http.MethodPost, "/api/login", loginHandlers.Login)

	return app
}
