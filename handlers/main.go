package handlers

import (
	"database/sql"
	"os"

	"github.com/vietthangc1/mini-web-golang/cache"
)

type BaseHandler struct{
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

var cacheInstance cache.CacheProducts = cache.CreateCache(os.Getenv("redisHost"), 0, 10 *1000000000) // db 0, expire 10s
