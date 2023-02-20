package repository

import "github.com/vietthangc1/mini-web-golang/models"


type CacheProducts interface {
	Set(key string, value models.Product) error
	Get(key string) (models.Product, error)
	Delete(key string) error
}