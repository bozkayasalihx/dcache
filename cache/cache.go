package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	Cacher
	data map[string][]byte
	mu   sync.RWMutex
}

func New() *Cache {
	return &Cache{
		data: make(map[string][]byte),
	}
}

func (c *Cache) Get(key []byte) ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.data[string(key)]
	if !ok {
		return []byte(""), nil
	}
	return val, nil
}

func (c *Cache) Set(key []byte, val []byte, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ttl > 0 {
		go func() {
			<-time.After(ttl)
			delete(c.data, string(key))
		}()
	}

	c.data[string(key)] = val
	return nil

}

func (c *Cache) Update(key []byte, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, ok := c.data[string(key)]
	if !ok {
		return fmt.Errorf("key not found %s", string(key))
	}

	c.data[string(key)] = val
	return nil
}

func (c *Cache) Delete(key []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, string(key))
	return nil
}
