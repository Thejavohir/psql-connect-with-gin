package redis

import (
	"context"
	"log"

	"psql/config"
	"psql/storage"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	rdb *redis.Client
	product *productRepo
}

func NewConnectionRedis(c *config.Config) (storage.CacheI, error) {
	var client = redis.NewClient(
		&redis.Options{
			Addr:     c.RedisHost + c.RedisPort,
			Password: c.PostgresPassword,
			DB:       c.RedisDB,
		},
	)

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Close()

	return &Cache{
		rdb: client,
	}, nil
}

func (c *Cache) Close() {
	c.rdb.Close()
}

func (c *Cache) Product() storage.ProductRepoCacheI {
	if c.product == nil {
		c.product = NewProductRepo(c.rdb)
	}

	return c.product
}