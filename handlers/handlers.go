package handlers

import (
	"github.com/vietthangc1/mini-web-golang/cache"
	"github.com/vietthangc1/mini-web-golang/modules"
)

type Handler struct {
	ProductRepo   modules.ProductRepository
	UserRepo      modules.UserRepository
	CacheInstance cache.CacheProducts
}

func NewHandler(productRepo modules.ProductRepository, userRepo modules.UserRepository, cacheInstance cache.CacheProducts) Handler {
	return Handler{
		ProductRepo:   productRepo,
		UserRepo:      userRepo,
		CacheInstance: cacheInstance,
	}
}
