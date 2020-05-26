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

// func main() {
// 	filePath := "tmp.db"

// 	imgBuf, err := ioutil.ReadFile("IMG_20190630_150511.jpg")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("imgBuf len", len(imgBuf))

// 	c := fastcache.New(128 * 1024 * 1024)
// 	defer func() {
// 		c.SaveToFile(filePath)
// 		c.Reset()
// 	}()

// 	if _, err := os.Stat(filePath); err == nil {
// 		c, err = fastcache.LoadFromFile(filePath)
// 	}

// 	fmt.Println("aaa:", string(c.Get(nil, []byte("aaa"))))

// 	c.Set([]byte("key"), []byte("value"))
// 	fmt.Println("key:", string(c.Get(nil, []byte("key"))))

// 	c.Set([]byte("aaa"), []byte("bbb"))
// 	fmt.Println("aaa:", string(c.Get(nil, []byte("aaa"))))

// 	c.Del([]byte("key"))

// 	var buf []byte
// 	if !c.Has([]byte("img")) {
// 		fmt.Println("!has img key, SetBig")
// 		c.SetBig([]byte("img"), imgBuf)
// 	}

// 	buf = c.GetBig(buf[:0], []byte("img"))
// 	if !bytes.Equal(buf, imgBuf) {
// 		fmt.Println("not Equal")
// 		return
// 	}
// 	fmt.Println("Equal", len(imgBuf), len(buf))
// }
