//go:build wireinject
//+build wireinject

package main

// import (
// 	"github.com/google/wire"
// 	"github.com/vietthangc1/mini-web-golang/app"
// 	"github.com/vietthangc1/mini-web-golang/handlers"
// 	"github.com/vietthangc1/mini-web-golang/models"
// 	"github.com/vietthangc1/mini-web-golang/repository/products"
// 	"github.com/vietthangc1/mini-web-golang/repository/redis"
// 	"github.com/vietthangc1/mini-web-golang/repository/users"
// )

// func InitializeApp() (app.App, error) {
// 	wire.Build(
// 		app.NewApp,
// 		app.NewRouter,
// 		handlers.NewHandler,
// 		products.NewProductService,
// 		users.NewUserService,
// 		models.ConnectDatabaseORM,
// 		redis.NewCacheInstance,
// 	)
// 	return app.App{}, nil
// }
