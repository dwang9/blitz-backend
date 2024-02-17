package dragonfly

import (
	"github.com/redis/go-redis/v9"

	"github.com/derek-test/blitz-backend/src/internal/cache"
)

type cacheClient struct {
	rdb redis.UniversalClient
}

func NewDragonflyCache(rdb redis.UniversalClient) cache.Cache {
	return &cacheClient{rdb: rdb}
}
