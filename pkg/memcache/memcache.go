package memcache

import (
	"context"
	"fmt"
	"sync"
)

type memCacheImpl[T any] struct {
	store map[string]T
	mu    sync.RWMutex
}

func NewMemCache[T any]() MemCache[T] {
	return &memCacheImpl[T]{
		store: make(map[string]T),
	}
}

func (c *memCacheImpl[T]) Set(_ context.Context, id string, val T) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.store[id] = val
	return nil
}

func (c *memCacheImpl[T]) Get(_ context.Context, id string) (T, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.store[id]
	if !ok {
		var zero T
		return zero, fmt.Errorf("data not found")
	}
	return val, nil
}

func (c *memCacheImpl[T]) Remove(_ context.Context, id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.store, id)
	return nil
}
