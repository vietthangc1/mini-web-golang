// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/vietthangc1/mini-web-golang/app"
	"github.com/vietthangc1/mini-web-golang/handlers"
	"github.com/vietthangc1/mini-web-golang/repository/mysql/db"
	"github.com/vietthangc1/mini-web-golang/repository/mysql/products"
	"github.com/vietthangc1/mini-web-golang/repository/mysql/users"
	"github.com/vietthangc1/mini-web-golang/repository/redis"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeApp() (app.App, error) {
	gormDB, err := db.ConnectDatabaseORM()
	if err != nil {
		return app.App{}, err
	}
	productRepo := products.NewProductRepo(gormDB)
	userRepo := users.NewUserRepo(gormDB)
	cacheProducts := redis.NewCacheInstance()
	handler := handlers.NewHandler(productRepo, userRepo, cacheProducts)
	engine := app.NewRouter(handler)
	appApp := app.NewApp(engine, handler)
	return appApp, nil
}
