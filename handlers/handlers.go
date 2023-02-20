package handlers

import (
	"github.com/vietthangc1/mini-web-golang/repository"
)

type Handler struct {
	ProductRepo   repository.ProductService
	UserRepo      repository.UserService
	CacheInstance repository.CacheProducts
}

func NewHandler(productServ repository.ProductService, userServ repository.UserService, cacheInstance repository.CacheProducts) Handler {
	return Handler{
		ProductRepo:   productServ,
		UserRepo:      userServ,
		CacheInstance: cacheInstance,
	}
}
