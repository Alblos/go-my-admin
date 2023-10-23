package cache

import (
	"context"
	"github.com/go-my-admin/server/logger"
	"github.com/redis/go-redis/v9"
)

type RedisConnection struct {
	Cnx *redis.Client
	Ctx context.Context
}

func (cache *RedisConnection) Connect(addr string, password string, db int) (err error) {
	cache.Cnx = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	cache.Ctx = context.Background()

	err = cache.Cnx.Set(cache.Ctx, "key", "value", 0).Err()
	if err != nil {
		logger.Error("Error connecting to internal cache: ", err)
		return err
	}

	err = cache.Cnx.Del(cache.Ctx, "key").Err()
	if err != nil {
		logger.Error("Error connecting to internal cache: ", err)
		return err
	}

	logger.Info("Connected to redis " + addr)
	return nil
}
