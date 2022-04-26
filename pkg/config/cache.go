package config

import (
	"gopssh/pkg/cache"
	"gopssh/pkg/file"
)

func SaveNewCache(configPath string, instances cache.Instances) error {
	cache := &cache.Cache{
		ConfigPath: configPath,
		ModTime:    file.MustGetFileModTime(configPath),
		Instances:  instances,
	}

	return cache.Save()
}
