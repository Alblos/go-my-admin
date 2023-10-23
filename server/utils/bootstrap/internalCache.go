package bootstrap

import (
	"github.com/go-my-admin/server/cache"
	"github.com/go-my-admin/server/logger"
)

func BootInternalCache(internalCacheObject *cache.RedisConnection) error {
	addr, password, db := cache.GetConnectionData()
	err := internalCacheObject.Connect(addr, password, db)
	if err != nil {
		logger.Error("Error connecting to internal cache: ", err)
		return err
	}

	return nil
}
