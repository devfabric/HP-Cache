package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var HPCacheConfig = &CacheConfig{
	MaxMem:  128 * 1024 * 1024,
	MapFile: filepath.Join("data", "HP-Cache.db"),
}

type CacheConfig struct {
	MaxMem  int    `json:"maxmem"`
	MapFile string `json:"mapfile"`
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func LoadHPCacheConfig(dir string) (*CacheConfig, error) {
	path := filepath.Join(dir, "./configs/cache.toml")
	filePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	config := new(CacheConfig)
	if CheckFileIsExist(filePath) { //文件存在
		if _, err := toml.DecodeFile(filePath, config); err != nil {
			return nil, err
		} else {
			HPCacheConfig = config
		}
	} else {
		configBuf := new(bytes.Buffer)
		if err := toml.NewEncoder(configBuf).Encode(HPCacheConfig); err != nil {
			return nil, err
		}
		err := ioutil.WriteFile(filePath, configBuf.Bytes(), 0666)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}
