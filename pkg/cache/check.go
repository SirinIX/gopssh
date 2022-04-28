package cache

import (
	"gopssh/log"
	"gopssh/pkg/file"
)

func GetInstancesByConfigPath(configPath string) Instances {
	cache := GetCacheByConfigPath(configPath)
	if cache == nil {
		return nil
	}

	if !file.IsPathExist(cache.InstancesPath) {
		return nil
	}

	if cache.IsConfigModTimeChanged() {
		return nil
	}

	instances, err := UnmarshalInstances(cache.InstancesPath)
	if err != nil {
		return nil
	}

	return instances
}

func GetCacheByConfigPath(cfgPath string) *Cache {
	cacheIndexPath := GetCacheIndexPath()
	if !file.IsPathExist(cacheIndexPath) {
		log.Warning("cache index file %v not found", cacheIndexPath)
		return nil
	}

	index, err := UnmarshalCacheIndex(cacheIndexPath)
	if err != nil {
		log.Warning("failed to unmarshal cache index, error: %s", err)
		return nil
	}

	cache := index.GetCacheByConfigPath(cfgPath)
	if cache == nil {
		log.Warning("cache file not found for config %v in cache index", cfgPath)
		return nil
	}

	return cache
}

func (c *Cache) IsConfigModTimeChanged() bool {
	aft, err := file.GetFileModTime(c.ConfigPath)
	if err != nil {
		log.Warning("failed to get file mod time for config %v, error: %s", c.ConfigPath, err)
		return true
	}

	if c.ModTime != aft {
		log.Warning("config %v mod time changed, old: %v, new: %v", c.ConfigPath, c.ModTime, aft)
		return true
	}

	return false
}
