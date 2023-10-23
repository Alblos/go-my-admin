package cache

import "os"

func GetConnectionData() (addr string, password string, db int) {
	return os.Getenv("REDIS_INTERNAL_CACHE_ADDR"), os.Getenv("REDIS_INTERNAL_CACHE_PASSWORD"), 0
}
