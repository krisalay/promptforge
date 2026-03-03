package memcache

import (
	"context"
	"sync"
	"testing"
)

func TestMemCache_SetAndGet(t *testing.T) {
	cache := NewMemCache[string]()
	ctx := context.Background()

	err := cache.Set(ctx, "key1", "value1")
	if err != nil {
		t.Fatalf("unexpected error on Set: %v", err)
	}

	val, err := cache.Get(ctx, "key1")
	if err != nil {
		t.Fatalf("unexpected error on Get: %v", err)
	}

	if val != "value1" {
		t.Fatalf("expected value 1, got %v", val)
	}
}

func TestMemCache_Get_NotFound(t *testing.T) {
	cache := NewMemCache[int]()
	ctx := context.Background()

	val, err := cache.Get(ctx, "missing")

	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if val != 0 {
		t.Fatalf("expected zero value (0), got %v", val)
	}
}

func TestMemCache_Remove_ExistingKey(t *testing.T) {
	cache := NewMemCache[string]()
	ctx := context.Background()

	_ = cache.Set(ctx, "key1", "value1")
	_ = cache.Remove(ctx, "key1")

	_, err := cache.Get(ctx, "key1")
	if err == nil {
		t.Fatal("expected error after removal, got nil")
	}
}

func TestMemCache_Remove_NonExistingKey(t *testing.T) {
	cache := NewMemCache[string]()
	ctx := context.Background()

	err := cache.Remove(ctx, "missing")
	if err != nil {
		t.Fatalf("unexpected error removing non-existing key: %v", err)
	}
}

func TestMemCache_OverwriteValue(t *testing.T) {
	cache := NewMemCache[string]()
	ctx := context.Background()

	_ = cache.Set(ctx, "key1", "value1")
	_ = cache.Set(ctx, "key1", "value2")

	val, err := cache.Get(ctx, "key1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if val != "value2" {
		t.Fatalf("expected value2, got %v", val)
	}
}

func TestMemCache_ConcurrentAccess(t *testing.T) {
	cache := NewMemCache[int]()
	ctx := context.Background()

	var wg sync.WaitGroup

	// Writers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = cache.Set(ctx, "key", i)
		}(i)
	}

	// Readers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, _ = cache.Get(ctx, "key")
		}()
	}

	wg.Wait()
}

type mockStruct struct {
	Name string
}

func TestMemCache_PointerType(t *testing.T) {
	cache := NewMemCache[*mockStruct]()
	ctx := context.Background()

	obj := &mockStruct{Name: "test"}
	_ = cache.Set(ctx, "id1", obj)

	val, err := cache.Get(ctx, "id1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if val.Name != "test" {
		t.Fatalf("expected test, got %v", val.Name)
	}
}
