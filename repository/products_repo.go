package repository

import (
	"github.com/vietthangc1/mini-web-golang/models"
)

type ProductService interface {
	GetProductByID(id uint) (models.Product, error)
	GetProducts(productFilter, propertisesFilter map[string]interface{}) ([]models.Product, error)
	AddProduct(newProduct models.Product) (models.Product, error)
	UpdateProduct(updateProduct models.Product, id uint) (models.Product, error)
	DeleteProduct(id uint) (models.Product, error)
}