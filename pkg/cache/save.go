package cache

import (
	"gopssh/log"
	"gopssh/pkg/file"
	"path/filepath"
)

func CreateOrUpdateCacheIndex(cache *Cache) error {
	cacheDir := GetCacheDir()
	if err := file.EnsureDirExist(cacheDir); err != nil {
		log.Warning("failed to ensure cache directory %v exist, error: %s", cacheDir, err)
		return err
	}

	// Create cache
	cacheIndexPath := GetCacheIndexPath()
	if !file.IsPathExist(cacheIndexPath) {
		index := &CacheIndex{cache}
		indexStr, err := index.ToYAML()
		if err != nil {
			return err
		}
		if err := file.SaveStringAsFile(cacheIndexPath, indexStr); err != nil {
			return err
		}

		return nil
	}

	// Update cache
	index, err := UnmarshalCacheIndex(cacheIndexPath)
	if err != nil {
		return err
	}
	if err := index.UpdateCache(cache); err != nil {
		return err
	}

	return nil
}

func (i CacheIndex) IsConfigExist(configPath string) int {
	for id, c := range i {
		if c.ConfigPath == configPath {
			return id
		}
	}

	return -1
}

func (i CacheIndex) UpdateCache(cache *Cache) error {
	if id := i.IsConfigExist(cache.ConfigPath); id != -1 {
		i[id] = cache
	} else {
		i = append(i, cache)
	}

	return i.Save()
}

func (i CacheIndex) Save() error {
	b, err := i.ToYAMLBytes()
	if err != nil {
		return err
	}

	cacheIndexPath := GetCacheIndexPath()
	if err := file.SaveBytesAsFile(cacheIndexPath, b); err != nil {
		return err
	}

	return nil
}

func (c *Cache) SaveInstances() error {
	cacheDir := filepath.Dir(c.InstancesPath)
	if err := file.EnsureDirExist(cacheDir); err != nil {
		return err
	}

	return c.Instances.Save(c.InstancesPath)
}

func (i Instances) Save(path string) error {
	b, err := i.ToYAMLBytes()
	if err != nil {
		return err
	}

	if err := file.SaveBytesAsFile(path, b); err != nil {
		return err
	}

	return nil
}
