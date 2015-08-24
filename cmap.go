package cmap

import "sync"

// CMap is a thread-safe map.
type CMap struct {
	items map[interface{}]interface{}
	mutex sync.RWMutex
}

// New creates and returns a new thread-safe map.
func New() *CMap {
	return &CMap{items: make(map[interface{}]interface{})}
}

// Get retrieves a value for a given key.
func (c *CMap) Get(key interface{}) (val interface{}, ok bool) {
	c.mutex.RLock()
	val, ok = c.items[key]
	c.mutex.RUnlock()
	return
}

// Set stores a value for a given key.
func (c *CMap) Set(key interface{}, val interface{}) {
	c.mutex.Lock()
	c.items[key] = val
	c.mutex.Unlock()
}

// Delete removes a value for a given key.
func (c *CMap) Delete(key interface{}) {
	c.mutex.Lock()
	delete(c.items, key)
	c.mutex.Unlock()
}
