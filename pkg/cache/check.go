package cache

import "gopssh/pkg/file"

func IsCacheFileExist(cfgPath string) (string, bool) {
	c := &Cache{
		ConfigPath: cfgPath,
	}

	cachePath := c.GetCacheFilePath()
	return cachePath, file.IsPathExist(cachePath)
}

func (c *Cache) IsConfigFileChanges() bool {
	bef := c.ModTime
	aft := file.MustGetFileModTime(c.ConfigPath)

	return aft != bef
}