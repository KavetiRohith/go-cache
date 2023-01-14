package cache

import (
	"fmt"
	"time"
)

type Cache struct {
	data map[string]string
}

func New() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, error) {
	val, ok := c.data[key]
	if !ok {
		return "", fmt.Errorf("key (%s) not found", key)
	}

	return val, nil

}

func (c *Cache) Has(key string) bool {
	_, isPresent := c.data[key]
	return isPresent
}

func (c *Cache) Set(key, val string) error {
	c.data[key] = val
	return nil
}

func (c *Cache) SetWithTTL(key, val string, ttl time.Duration) error {
	// TODO: Implement expiry mechanism
	c.data[key] = val
	return nil
}

func (c *Cache) Delete(key string) error {
	delete(c.data, key)
	return nil
}
