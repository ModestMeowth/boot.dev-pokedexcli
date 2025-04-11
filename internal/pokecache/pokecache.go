package pokecache

import (
    "sync"
    "time"
)

type Cache struct {
    interval time.Duration
    mutex *sync.Mutex
    entries map[string]cacheEntry
}

type cacheEntry struct {
    createdAt time.Time
    val []byte
}

func NewCache(interval time.Duration) Cache {
    cache := Cache{
        interval: interval,
        mutex: &sync.Mutex{},
        entries: map[string]cacheEntry{},
    }
    go cache.reapLoop()
    return cache
}

func (c *Cache) Add(key string, val []byte) error {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    c.entries[key] = cacheEntry{
        createdAt: time.Time{},
        val: val,
    }
    return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mutex.Lock()
    defer c.mutex.Unlock()
    entry, ok := c.entries[key]
    if !ok {
        return []byte{}, false
    } else {
        return entry.val, true
    }
}

func (c *Cache) reapLoop() {
    ticker := time.NewTicker(c.interval)
    defer ticker.Stop()

    for cTime := range ticker.C {
        for key, entry := range c.entries {
            cutoff := cTime.Add(-c.interval)
            if entry.createdAt.Before(cutoff) {
                c.mutex.Lock()
                delete(c.entries, key)
                c.mutex.Unlock()
            }
        }
    }
}
