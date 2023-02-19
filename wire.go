//+build wireinject

package main

// import (
// 	"github.com/google/wire"
// 	"github.com/vietthangc1/mini-web-golang/app"
// 	"github.com/vietthangc1/mini-web-golang/cache"
// 	"github.com/vietthangc1/mini-web-golang/handlers"
// 	"github.com/vietthangc1/mini-web-golang/models"
// 	"github.com/vietthangc1/mini-web-golang/modules"
// )

// func InitializeApp() (app.App, error) {
// 	wire.Build(
// 		app.NewApp,
// 		app.NewRouter,
// 		handlers.NewHandler,
// 		modules.NewProductRepository,
// 		modules.NewUserRepository,
// 		models.ConnectDatabaseORM,
// 		cache.NewCacheInstance,
// 	)
// 	return app.App{}, nil
// }
