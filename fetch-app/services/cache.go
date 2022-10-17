package services

import "github.com/patrickmn/go-cache"

func saveCache(c *cache.Cache, key string, value interface{}) {
	c.Set(key, value, cache.DefaultExpiration)
}

func getCache(c *cache.Cache, key string) (interface{}, bool) {
	data, exist := c.Get("foo")
	return data, exist
}
