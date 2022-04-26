package cache

import "gopssh/pkg/file"

func IsCacheFileExist(cfgPath string) bool {
	c := &Cache{
		ConfigPath: cfgPath,
	}

	return file.IsFileExist(c.GetCacheFilePath()) 
}