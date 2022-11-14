package redis_cache

import (
	"calculator-ddd/pkg/utils"
	"github.com/go-redis/redis/v8"
)

func InitRedisCache(config utils.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.Db,       // use default DB
	})

	return rdb, nil
}
