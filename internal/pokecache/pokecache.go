package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
type PokeCache struct {
	CacheMap map[string]cacheEntry
	Mut      sync.Mutex
}

func NewCache(interval time.Duration) PokeCache {
	c := PokeCache{CacheMap: map[string]cacheEntry{}}
	c.reapLoop(interval)
	return c
}

func (c *PokeCache) Add(key string, val []byte) {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	c.CacheMap[key] = cacheEntry{createdAt: time.Now(), val: val}

}
func (c *PokeCache) Get(key string) ([]byte, bool) {
	c.Mut.Lock()
	cacheEntry, ok := c.CacheMap[key]
	defer c.Mut.Unlock()
	if !ok {
		return nil, false
	}
	return cacheEntry.val, true
}
func (c *PokeCache) reapLoop(interval time.Duration) {

	go func(ca *PokeCache) {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			<-ticker.C
			for k, v := range ca.CacheMap {
				if time.Since(v.createdAt) > interval {
					ca.Mut.Lock()
					delete(ca.CacheMap, k)
					ca.Mut.Unlock()
				}
			}
		}
	}(c)
}
