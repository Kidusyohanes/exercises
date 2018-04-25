package main

import (
	"log"
	"sync"
	"time"
)

//ttlentry is an entry in a TTLCache.
type ttlentry struct {
	expires time.Time
	value   interface{}
}

//TTLCache is an in-memory time-to-live cache that is safe for concurrent use.
type TTLCache struct {
	entries map[string]*ttlentry
	mx      sync.RWMutex
}

//NewTTLCache constructs a new TTLCache.
func NewTTLCache() *TTLCache {
	c := &TTLCache{
		entries: map[string]*ttlentry{},
	}
	go c.janitor()
	return c
}

//Set sets a key and value in the cache. The key must be a string,
//but the value can be of any time. The ttl is how long the entry
//should remain the cache.
func (c *TTLCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.entries[key] = &ttlentry{
		expires: time.Now().Add(ttl),
		value:   value,
	}
}

//Get gets the value associated with the key.
func (c *TTLCache) Get(key string) interface{} {
	c.mx.RLock()
	defer c.mx.RUnlock()

	e, found := c.entries[key]
	if !found {
		return nil
	}
	return e.value
}

func (c *TTLCache) Purge() {
	c.mx.Lock()
	defer c.mx.Unlock()

	for k, e := range c.entries {
		if e.expires.Before(time.Now()) {
			delete(c.entries, k)
		}
	}
}

func (c *TTLCache) janitor() {
	for {
		time.Sleep(time.Second * 5)
		log.Printf("RUNNING PURGE")
		c.Purge()
		log.Printf("COMPLETED PURGE")
	}
}
