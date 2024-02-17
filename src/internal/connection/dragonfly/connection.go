package dragonfly

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	redistrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/redis/go-redis.v9"

	"github.com/derek-test/blitz-backend/src/internal/config"
)

func GetDragonflyConnection(cfg *config.Config) redis.UniversalClient {
	return redistrace.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.CacheConfig.Host, cfg.CacheConfig.Port),
		Password: cfg.CacheConfig.Password,
		DB:       0,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Info().Msg("connected to the cache")
			return nil
		},
		ContextTimeoutEnabled: true,
	})
}
