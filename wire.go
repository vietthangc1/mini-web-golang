package main

import (
	"github.com/google/wire"
	"github.com/vietthangc1/mini-web-golang/app"
	"github.com/vietthangc1/mini-web-golang/cache"
	"github.com/vietthangc1/mini-web-golang/handlers"
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/modules"
)

func InitializeApp() app.App {
	wire.Build(
		app.NewApp,
		app.NewCacheInstance,
		models.ConnectDatabaseORM,
		handlers.NewHandler,
		modules.NewProductRepository,
		modules.NewUserRepository,
		cache.NewCache,
	)
	return app.App{}
}
