package store

import (
	"sync"
	"time"
)

var cacheInstance *Cache
var once = sync.Once{}

// GetCacheInstance returns a singleton instance of the cache
func GetCacheInstance() *Cache {
	once.Do(func() {
		if cacheInstance == nil {
			cacheInstance = &Cache{
				data: make(map[string]Item),
			}
		}
	})
	return cacheInstance
}

type Cache struct {
	lock sync.Mutex
	data map[string]Item
}

type Item struct {
	Key   string
	Value string
	TTL   int64
}

func (c *Cache) Set(key string, value Item) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = value
	return nil
}

func (c *Cache) Del(key string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.data, key)
	return nil
}

func (c *Cache) Get(key string) (Item, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	item, found := c.data[key]
	if found && item.Expired() {
		delete(c.data, key)
		found = false
	}
	return item, found
}

func (it *Item) Expired() bool {
	return it.TTL > 0 && time.Now().UnixNano() > it.TTL
}
