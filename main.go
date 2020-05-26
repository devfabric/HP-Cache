package main

import (
	"fmt"

	"github.com/devfabric/HP-Cache/cache"
	"github.com/devfabric/HP-Cache/config"
)

func main() {
	cacheConfig, err := config.LoadHPCacheConfig("./")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(cacheConfig)

	myCache, err := cache.NewCache()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err = myCache.ToSaveFile()
		if err != nil {
			fmt.Println(err.Error())
		}
		myCache.ResetMem()
	}()

	myCache.Set([]byte("key1"), []byte("value1"))
	if myCache.Has([]byte("key1")) {
		fmt.Println(string(myCache.Get([]byte("key1"))))
	}

	myCache.ResetMem()
	if !myCache.Has([]byte("key1")) {
		fmt.Println("key1 not find")
	}

	myCache.UpdateStats()
}
