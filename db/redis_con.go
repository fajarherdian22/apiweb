package db

import (
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func NewRedisClient(host, password string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
}

func GetRedisClient() *redis.Client {
	return rdb
}
