package repository

import (
	"net/url"

	"github.com/vietthangc1/mini-web-golang/models"
)

type ProductRepo interface {
	GetProductByID(id uint) (models.Product, error)
	GetProducts(filter url.Values) ([]models.Product, error)
	AddProduct(newProduct models.Product) (models.Product, error)
	UpdateProduct(updateProduct models.Product, id uint) (models.Product, error)
	DeleteProduct(id uint) (models.Product, error)
}