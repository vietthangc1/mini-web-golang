package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vietthangc1/mini-web-golang/models"
)

var ctx = context.Background()

type CacheInfo struct {
	Host   string
	DB     int
	Expire time.Duration
}

type CacheProducts interface {
	Set(key string, value models.Product) error
	Get(key string) (models.Product, error)
	Delete(key string) error
}

func NewCache(host string, db int, expireTime time.Duration) *CacheInfo {
	return &CacheInfo{
		Host:   host,
		DB:     db,
		Expire: expireTime,
	}
}

func NewCacheInstance() CacheProducts {
	return NewCache(os.Getenv("REDISHOST"), 0, 10*1000000000) // db 0, expire 10s
}

func (c *CacheInfo) getClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Host,
		Password: "",   // no password set
		DB:       c.DB, // use default DB
	})
	return rdb
}

func (c *CacheInfo) Set(key string, value models.Product) error {
	rdb := c.getClient()

	out, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = rdb.Set(ctx, key, string(out), c.Expire).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *CacheInfo) Get(key string) (models.Product, error) {
	rdb := c.getClient()
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return models.Product{}, err
	}
	out := models.Product{}
	json.Unmarshal([]byte(val), &out)
	return out, nil
}

func (c *CacheInfo) Delete(key string) error {
	rdb := c.getClient()
	searchPattern := key

	if len(os.Args) > 1 {
		searchPattern = os.Args[1]
	}

	var foundedRecordCount int = 0
	iter := rdb.Scan(ctx, 0, searchPattern, 0).Iterator()
	fmt.Printf("YOUR SEARCH PATTERN= %s\n", searchPattern)
	for iter.Next(ctx) {
		fmt.Printf("Deleted= %s\n", iter.Val())
		rdb.Del(ctx, iter.Val())
		foundedRecordCount++
	}
	if err := iter.Err(); err != nil {
		return err
	}
	fmt.Printf("Deleted Count %d\n", foundedRecordCount)
	return nil
}
