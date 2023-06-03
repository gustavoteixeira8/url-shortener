package config

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/gustavoteixeira8/url-shortener/src/utils"
)

func CacheConfig() *redis.Options {
	redisPort := utils.GetEnv("REDIS_PORT")
	redisHost := utils.GetEnv("REDIS_HOST")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	return &redis.Options{
		Addr:     redisAddr,
		Password: utils.GetEnv("REDIS_PASSWORD"),
		DB:       0,
	}
}
