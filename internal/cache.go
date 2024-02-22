package internal

import (
	"fmt"
	"log"
	"sync"
)

var Cache Cacher

type Cacher interface {
	Has(string) bool
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

func (c *InMemoryCache) Has(key string) bool {
	c.muData.RLock()
	defer c.muData.RUnlock()

	_, ok := c.data[key]
	return ok
}

func (c *InMemoryCache) Get(key string) (*any, bool) {
	c.muData.RLock()
	defer c.muData.RUnlock()

	val, ok := c.data[key]
	return val, ok
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

	log.Printf("purging cache for %s", key)

	delete(c.data, key)
	return nil
}

func (c *InMemoryCache) Purge() {
	c.muData.Lock()
	defer c.muData.Unlock()

	c.data = make(map[string]*any)
}
