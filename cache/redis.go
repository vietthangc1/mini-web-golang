package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vietthangc1/mini-web-golang/models"
)

var ctx = context.Background()

type CacheInfo struct {
    Host string
    DB int
    Expire time.Duration
}

type CacheFunction interface{
    Set(key string, value models.Product)
    Get(key string)
}

func CreateCache(host string, db int, expireTime time.Duration) (*CacheInfo) {
    return &CacheInfo{
        Host: host,
        DB: db,
        Expire: expireTime,
    }
}

func (c *CacheInfo) getClient() (*redis.Client) {
    rdb := redis.NewClient(&redis.Options{
        Addr:     c.Host,
        Password: "", // no password set
        DB:       c.DB,  // use default DB
    })
    return rdb
}

func (c *CacheInfo) Set(key string, value models.Product) (error) {
    rdb := c.getClient()

    out, err := json.Marshal(value)
	if (err != nil) {
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