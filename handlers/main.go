package handlers

import (
	"os"

	"github.com/vietthangc1/mini-web-golang/cache"
	"gorm.io/gorm"
)

type BaseHandler struct{
	db *gorm.DB
}

func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

var cacheInstance cache.CacheProducts = cache.CreateCache(os.Getenv("redisHost"), 0, 10 *1000000000) // db 0, expire 10s
