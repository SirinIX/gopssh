package cache

import (
	"fmt"
	"gopssh/pkg/file"
	"path/filepath"
)

const (
	CacheDirName    = ".gopssh/cache"
	CacheFilePrefix = ".cache_"
)

type Cache struct {
	ConfigPath  string    `json:"config_path" yaml:"config_path"`
	TimeVersion string    `json:"time_version" yaml:"time_version"`
	Instances   Instances `json:"instances" yaml:"instances"`
}

type Instances []*Instance

type Instance struct {
	Address  *Address          `json:"address" yaml:"address"`
	Username string            `json:"username" yaml:"username"`
	Password string            `json:"password" yaml:"password"`
	Labels   map[string]string `json:"labels" yaml:"labels"`
}

type Address struct {
	Ip   string `json:"ip" yaml:"ip"`
	Port int    `json:"port" yaml:"port"`
}

func (a *Address) String() string {
	return fmt.Sprintf("%s:%d", a.Ip, a.Port)
}

func GetCacheDir() string {
	return filepath.Join(file.MustGetUserHome(), CacheDirName)
}

func (c *Cache) GetCacheFileName() string {
	return CacheFilePrefix + filepath.Base(c.ConfigPath)
}

func (c *Cache) GetCacheFilePath() string {
	return filepath.Join(GetCacheDir(), c.GetCacheFileName())
}

func (c *Cache) Save() error {
	cacheStr, err := c.ToYAML()
	if err != nil {
		return err
	}

	// Ensure cache directory exist
	cacheDir := GetCacheDir()
	if err := file.EnsureDirExist(cacheDir); err != nil {
		return err
	}

	// Save cache file
	cacheFilePath := filepath.Join(cacheDir, c.GetCacheFileName())
	return file.SaveStringAsFile(cacheFilePath, cacheStr)
}
