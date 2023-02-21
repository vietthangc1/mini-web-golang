//go:build wireinject
//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/vietthangc1/mini-web-golang/app"
	"github.com/vietthangc1/mini-web-golang/handlers"
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/repository/mysql/products"
	"github.com/vietthangc1/mini-web-golang/repository/redis"
	"github.com/vietthangc1/mini-web-golang/repository/mysql/users"
)

func InitializeApp() (app.App, error) {
	wire.Build(
		app.NewApp,
		app.NewRouter,
		handlers.NewHandler,
		products.NewProductRepo,
		users.NewUserRepo,
		models.ConnectDatabaseORM,
		redis.NewCacheInstance,
	)
	return app.App{}, nil
}
