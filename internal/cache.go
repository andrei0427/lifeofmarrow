package internal

import (
	"fmt"
	"sync"
)

var Cache Cacher

type Cacher interface {
	Get(string) (*any, bool)
	Set(string, any) error
	Invalidate(string) error
	Purge()
}

type InMemoryCache struct {
	Cacher
	muData sync.RWMutex
	data   map[string]*any
}

func NewCache() Cacher {
	return &InMemoryCache{
		data:   make(map[string]*any),
		muData: sync.RWMutex{},
	}
}

func (c *InMemoryCache) Get(key string) (*any, bool) {
	c.muData.RLock()
	defer c.muData.RUnlock()

	val, ok := c.data[key]
	return val, ok && val != nil
}

func (c *InMemoryCache) Set(key string, val any) error {
	c.muData.Lock()
	defer c.muData.Unlock()

	if _, exists := c.data[key]; exists {
		return fmt.Errorf("cache key %s already exists", key)
	}

	c.data[key] = &val
	return nil
}

func (c *InMemoryCache) Invalidate(key string) error {
	c.muData.Lock()
	defer c.muData.Unlock()

	if _, exists := c.data[key]; !exists {
		return fmt.Errorf("cache key %s doesn't exist", key)
	}

	c.data[key] = nil
	return nil
}

func (c *InMemoryCache) Purge() {
	c.muData.Lock()
	defer c.muData.Unlock()

	c.data = make(map[string]*any)
}
