package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func New() *Cache {
	return &Cache{
		data: make(map[string]string),
		mu:   sync.RWMutex{},
	}
}

func (c *Cache) Get(key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.data[key]
	if !ok {
		return "", fmt.Errorf("key (%s) not found", key)
	}

	return val, nil

}

func (c *Cache) Has(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	_, isPresent := c.data[key]
	return isPresent
}

func (c *Cache) Set(key, val string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = val

	return nil
}

func (c *Cache) SetWithTTL(key, val string, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	go func(c *Cache, key string, ttl time.Duration) {
		<-time.After(ttl)

		c.mu.Lock()
		defer c.mu.Unlock()

		delete(c.data, key)
		log.Printf("key %s with val %s expired\n", key, val)
	}(c, key, ttl)

	c.data[key] = val
	return nil
}

func (c *Cache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
	return nil
}
