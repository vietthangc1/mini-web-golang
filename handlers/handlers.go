package handlers

import (
	"github.com/vietthangc1/mini-web-golang/modules"
)

type Handler struct {
	ProductRepo modules.ProductRepository
	UserRepo    modules.UserRepository
}

func NewHandler(productRepo modules.ProductRepository, userRepo modules.UserRepository) Handler {
	return Handler{
		ProductRepo: productRepo,
		UserRepo:    userRepo,
	}
}
