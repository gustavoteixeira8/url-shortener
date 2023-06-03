package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	"github.com/gustavoteixeira8/url-shortener/src/config"
	"github.com/sirupsen/logrus"
)

var redisClient *redis.Client

func SetupCache() error {
	if redisClient == nil {
		redisClient = redis.NewClient(config.CacheConfig())
	}

	err := redisClient.Ping().Err()
	if err != nil {
		return err
	}

	logrus.Info("Cache connected")

	return nil
}

type AppCache[T any] struct {
	client *redis.Client
}

func (c AppCache[T]) Set(key string, val T, expiration time.Duration) error {
	valStr, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return c.client.Set(key, valStr, expiration).Err()
}

func (c AppCache[T]) Get(key string) (*T, error) {
	rresult := c.client.Get(key)

	if rresult.Err() != nil {
		return nil, rresult.Err()
	}

	valBytes, err := rresult.Bytes()

	if err != nil {
		return nil, err
	}

	val := new(T)

	err = json.Unmarshal(valBytes, val)
	if err != nil {
		return nil, err
	}

	return val, nil
}

func (c AppCache[T]) Delete(key ...string) error {
	return c.client.Del(key...).Err()
}

func NewAppCache[T any]() *AppCache[T] {
	return &AppCache[T]{client: redisClient}
}
