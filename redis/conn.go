package redis

import (
	"fmt"

	"github.com/coderc/onlinedisk-util/config"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func Init(config *config.RedisConfigStruct) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.DB,
	})
}

func GetConn() *redis.Client {
	return rdb
}
