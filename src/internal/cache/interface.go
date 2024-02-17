package cache

import (
	"context"
)

//go:generate mockgen -source interface.go -destination mocks/mock.go -mock_names Cache=MockCache
type Cache interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	Incr(ctx context.Context, key string) error
	Decr(ctx context.Context, key string) error
}
