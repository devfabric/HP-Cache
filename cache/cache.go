package cache

import (
	"fmt"
	"os"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/devfabric/HP-Cache/config"
)

type BigCache struct {
	Cache *fastcache.Cache
}

func NewCache() (*BigCache, error) {
	cache := &BigCache{
		Cache: fastcache.New(config.HPCacheConfig.MaxMem),
	}

	if _, err := os.Stat(config.HPCacheConfig.MapFile); err == nil {
		cache.Cache, err = fastcache.LoadFromFile(config.HPCacheConfig.MapFile)
		if err != nil {
			return nil, err
		}
	}

	return cache, nil
}

func (c *BigCache) ToSaveFile() error {
	if c != nil && c.Cache != nil {
		return c.Cache.SaveToFile(config.HPCacheConfig.MapFile)
	}
	return nil
}

func (c *BigCache) ResetMem() {
	if c != nil && c.Cache != nil {
		c.Cache.Reset()
	}
}

func (c *BigCache) Set(key, val []byte) {
	if c != nil && c.Cache != nil {
		c.Cache.Set(key, val)
	}
}

func (c *BigCache) Get(key []byte) []byte {
	if c != nil && c.Cache != nil {
		return c.Cache.Get(nil, key)
	}
	return nil
}

func (c *BigCache) Has(key []byte) bool {
	if c != nil && c.Cache != nil {
		return c.Cache.Has(key)
	}
	return false
}

func (c *BigCache) Del(key []byte) {
	if c != nil && c.Cache != nil {
		c.Cache.Del(key)
	}
}

func (c *BigCache) SetBig(key, bigVal []byte) {
	if c != nil && c.Cache != nil {
		c.Cache.SetBig(key, bigVal)
	}
}

func (c *BigCache) GetBig(buffer, key []byte) []byte {
	if c != nil && c.Cache != nil {
		c.Cache.GetBig(buffer[:0], key)
		return buffer
	}
	return nil
}

func (c *BigCache) UpdateStats() {
	if c != nil && c.Cache != nil {
		var stats fastcache.Stats
		c.Cache.UpdateStats(&stats)
		fmt.Println(stats)
	}
}
