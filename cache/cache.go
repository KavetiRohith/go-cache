package cache

import (
	"fmt"
	"log"
	"time"
)

type obj struct {
	value     string
	expiresAt int64
}

func newObj(value string, duration int64) *obj {
	var expiresAt int64 = -1
	if duration > 0 {
		expiresAt = time.Now().Unix() + duration
	}

	return &obj{
		value:     value,
		expiresAt: expiresAt,
	}
}

type Cache struct {
	data map[string]*obj
}

func New() *Cache {
	return &Cache{
		data: make(map[string]*obj),
	}
}

func (c *Cache) Get(key string) (string, error) {
	obj, ok := c.data[key]
	if !ok {
		return "", fmt.Errorf("key (%s) not found", key)
	}

	// passive deletion of expired keys when accessed
	if obj.expiresAt != -1 && obj.expiresAt <= time.Now().Unix() {
		delete(c.data, key)
		return "", fmt.Errorf("key (%s) not found", key)
	}

	return obj.value, nil
}

func (c *Cache) Has(key string) bool {
	_, isPresent := c.data[key]
	return isPresent
}

func (c *Cache) Set(key, val string) error {
	c.data[key] = newObj(val, -1)
	return nil
}

func (c *Cache) SetWithTTL(key, val string, ttl int64) error {
	c.data[key] = newObj(val, ttl)
	return nil
}

func (c *Cache) Delete(key string) error {
	delete(c.data, key)
	return nil
}

func (c *Cache) expireSample() float32 {
	var limit int = 20
	var expiredCount int = 0

	// assuming iteration of golang hash table in randomized
	for key, obj := range c.data {
		if obj.expiresAt != -1 {
			limit--
			// if the key is expired
			if obj.expiresAt <= time.Now().Unix() {
				delete(c.data, key)
				expiredCount++
			}
		}

		// once we iterated to 20 keys that have some expiration set
		// we break the loop
		if limit == 0 {
			break
		}
	}

	return float32(expiredCount) / float32(20.0)
}

// Deletes all the expired keys - the active way
// Sampling approach: https://redis.io/commands/expire/
func (c *Cache) DeleteExpiredKeys() {
	for {
		frac := c.expireSample()
		// if the sample had less than 25% keys expired
		// we break the loop.
		if frac < 0.25 {
			break
		}
	}
	log.Println("deleted the expired but undeleted keys. total keys", len(c.data))
}
