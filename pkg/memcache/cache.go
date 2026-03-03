package memcache

import "context"

type MemCache[T any] interface {
	Set(ctx context.Context, id string, val T) error
	Get(ctx context.Context, id string) (T, error)
	Remove(ctx context.Context, id string) error
}
