package handlers

import (
	"github.com/vietthangc1/mini-web-golang/repository"
)

type Handler struct {
	ProductRepo   repository.ProductRepo
	UserRepo      repository.UserRepo
	CacheInstance repository.CacheProducts
}

func NewHandler(productServ repository.ProductRepo, userServ repository.UserRepo, cacheInstance repository.CacheProducts) Handler {
	return Handler{
		ProductRepo:   productServ,
		UserRepo:      userServ,
		CacheInstance: cacheInstance,
	}
}
