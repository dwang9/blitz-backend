package dragonfly

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (c *cacheClient) Set(ctx context.Context, key string, value interface{}) error {
	err := c.rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to set key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (c *cacheClient) Del(ctx context.Context, key string) error {
	err := c.rdb.Del(ctx, key).Err()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to del key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) Incr(ctx context.Context, key string) error {
	err := c.rdb.Incr(ctx, key).Err()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to incr key in the cache")
		return err
	}

	return nil
}

func (c *cacheClient) Decr(ctx context.Context, key string) error {
	err := c.rdb.Decr(ctx, key).Err()
	if err != nil {
		log.
			Error().
			Err(err).
			Str("key", key).
			Msg("unable to decr key in the cache")
		return err
	}

	return nil
}
