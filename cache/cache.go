package main

import (
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
}

//NewTTLCache constructs a new TTLCache.
func NewTTLCache() *TTLCache {
	c := &TTLCache{
		entries: map[string]*ttlentry{},
	}
	return c
}

//Set sets a key and value in the cache. The key must be a string,
//but the value can be of any time. The ttl is how long the entry
//should remain the cache.
func (c *TTLCache) Set(key string, value interface{}, ttl time.Duration) {
	c.entries[key] = &ttlentry{
		expires: time.Now().Add(ttl),
		value:   value,
	}
}

//Get gets the value associated with the key.
func (c *TTLCache) Get(key string) interface{} {
	e, found := c.entries[key]
	if !found {
		return nil
	}
	return e.value
}
