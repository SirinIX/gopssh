package config

import (
	"gopssh/pkg/cache"
	"gopssh/pkg/file"
)

func SaveNewCache(configPath string, instances cache.Instances) error {
	ch := &cache.Cache{
		InstancesPath:  cache.GenerateCacheFilePathByConfig(configPath),
		ConfigPath: configPath,
		ModTime:    file.MustGetFileModTime(configPath),
		Instances:  instances,
	}

	if err := cache.CreateOrUpdateCacheIndex(ch); err != nil {
		return err
	}

	return ch.Instances.Save(ch.InstancesPath)
}
