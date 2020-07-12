package v201809

import (
	sha256 "crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"strings"
)

// cache

type callCache struct {
	Items map[string][]byte
}

var (
	cache_DIR     = ""
	cache_ENABLED = false
	cache         *callCache
)

func (c *callCache) Set(k []string, v []byte) {
	hashBuffer := sha256.Sum256([]byte(strings.Join(k, "-")))
	key := hex.EncodeToString(hashBuffer[:])
	c.Items[key] = v
}

func (c *callCache) Get(k []string) ([]byte, bool) {
	hashBuffer := sha256.Sum256([]byte(strings.Join(k, "-")))
	key := hex.EncodeToString(hashBuffer[:])
	if v, ok := c.Items[key]; ok {
		return v, ok
	} else {
		d, err := ioutil.ReadFile(cache_DIR + key)
		if err != nil {
			return []byte{}, false
		}
		return d, true
	}
}

func InitCache(dir string) {
	cache_ENABLED = true
	cache_DIR = dir
	cache = &callCache{
		Items: map[string][]byte{},
	}
}

func ResumeCache() {
	cache_ENABLED = true
}

func PauseCache() {
	cache_ENABLED = false
}

// call this method when you dont need memory cache
func SaveCache() error {
	for k, v := range cache.Items {
		err := ioutil.WriteFile(cache_DIR+k, []byte(v), 0777)
		if err != nil {
			return err
		}
	}
	// clear for memory save
	cache.Items = map[string][]byte{}
	return nil
}
