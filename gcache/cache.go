// Package gcache
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2023/2/10
 */
package gcache

import "time"
import "github.com/patrickmn/go-cache"

type Config struct {
	Enabled   bool          `json:"enabled" yaml:"enabled"`
	Timeout   time.Duration `json:"timeout" yaml:"timeout"`
	ClearTime time.Duration `json:"clear_time" yaml:"clear_time"`
}

var _cache *GCache

type GCache struct {
	cache *cache.Cache
}

func defaultCache() *GCache {
	// 设置超时时间和清理时间
	c := cache.New(10*time.Minute, 24*time.Hour)
	return &GCache{cache: c}
}

func NewCache(cfg *Config) *GCache {
	// 设置超时时间和清理时间
	if cfg == nil {
		_cache = defaultCache()
		return _cache
	}
	c := cache.New(cfg.Timeout, cfg.ClearTime)
	_cache = &GCache{cache: c}
	return _cache
}

func (gc *GCache) Set(key string, data interface{}, expire time.Duration) {
	gc.cache.Set(key, data, expire)
}

func (gc *GCache) Delete(key string) {
	gc.cache.Delete(key)
}

func (gc *GCache) Add(key string, data interface{}, expire time.Duration) {
	gc.cache.Add(key, data, expire)
}

func (gc *GCache) Get(key string) (interface{}, bool) {
	return gc.cache.Get(key)
}

func (gc *GCache) GetString(key string) (string, bool) {
	data, isExist := gc.cache.Get(key)
	if isExist {
		return data.(string), true
	}
	return "", false
}

func Get() *GCache {
	if _cache != nil {
		return _cache
	}
	_cache = defaultCache()
	return _cache
}
