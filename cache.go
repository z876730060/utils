package utils

import (
	"sync"
	"time"
)

var (
	Cache func() *CacheMap
)

func init() {
	Cache = sync.OnceValue(NewCacheMap)
}

type (
	handle interface {
		Set(key string, value any)
		SetDuration(key string, value any, duration time.Duration)
		Get(key string) (any, bool)
		Exist(key string) bool
		check()
		Delete(key string)
		Range(func(key string, value any) bool)
	}
	CacheMap struct {
		m  map[string]*cacheBean
		mu sync.Mutex
	}
	cacheBean struct {
		value    any
		duration *time.Duration
		time     time.Time
	}
)

func NewCacheMap() *CacheMap {
	c := &CacheMap{
		m: make(map[string]*cacheBean),
	}
	go c.check()
	return c
}

func (c *CacheMap) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = &cacheBean{
		value:    value,
		duration: nil,
		time:     time.Now(),
	}
}

func (c *CacheMap) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	bean, ok := c.m[key]
	if !ok {
		return nil, ok
	}
	return bean.value, true
}

func (c *CacheMap) SetDuration(key string, value any, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = &cacheBean{
		value:    value,
		duration: &duration,
		time:     time.Now(),
	}
}

func (c *CacheMap) check() {
	for {
		c.mu.Lock()
		for key, bean := range c.m {
			if bean.duration != nil && bean.time.Add(*bean.duration).Before(time.Now()) {
				delete(c.m, key)
			}
		}
		c.mu.Unlock()
		time.Sleep(1 * time.Second) // 1分钟检查一次
	}
}

func (c *CacheMap) Exist(key string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.m[key]
	return ok
}

func (c *CacheMap) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}

func (c *CacheMap) Range(rangeFunc func(key string, value any) bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, bean := range c.m {
		if !rangeFunc(key, bean.value) {
			break
		}
	}
}
