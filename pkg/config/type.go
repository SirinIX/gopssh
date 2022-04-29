package config

import (
	"path/filepath"

	"gopssh/log"
	"gopssh/pkg/cache"
	"gopssh/pkg/file"
)

const (
	DefaultConfigFileName = "inventory.yaml"

	GopsshDirName = ".gopssh"
)

type Config struct {
	Global *Global `json:"global" yaml:"global"`
	Groups Groups  `json:"groups" yaml:"groups"`
}

type Global struct {
	Port     int               `json:"port" yaml:"port"`
	Username string            `json:"username" yaml:"username"`
	Password string            `json:"password" yaml:"password"`
	// Labels   map[string]string `json:"labels" yaml:"labels"`
}

type Groups []*Group

type Group struct {
	Ips      []string          `json:"ips" yaml:"ips"`
	Port     int               `json:"port" yaml:"port"`
	Username string            `json:"username" yaml:"username"`
	Password string            `json:"password" yaml:"password"`
	Labels   map[string]string `json:"labels" yaml:"labels"`
}

func NewGlobal() *Global {
	return &Global{
		Port:     22,
		Username: "root",
	}
}

func NewConfig() *Config {
	return &Config{
		Global: NewGlobal(),
	}
}

func GetDefaultConfigFilePath() string {
	return filepath.Join(file.MustGetUserHome(), GopsshDirName, DefaultConfigFileName)
}

func IsDefaultConfigFileExist() bool {
	return file.IsPathExist(GetDefaultConfigFilePath())
}

func ConfigFileToInstances(path string, withoutCache bool) (cache.Instances, error) {
	// Check file
	cfgPath := path
	if path == "" {
		defaultCfgPath := GetDefaultConfigFilePath()
		if err := file.IsPathExistE(defaultCfgPath); err != nil {
			return nil, err
		}
		cfgPath = defaultCfgPath
	}
	if err := file.IsPathExistE(cfgPath); err != nil {
		return nil, err
	}

	// Use cache
	if !withoutCache {
		log.Debug("use cache for config file %v", cfgPath)
		cacheInstances := cache.GetInstancesByConfigPath(cfgPath)
		if cacheInstances != nil {
			return cacheInstances, nil
		}
	}

	// Use config directly
	cfg, err := UnmarshalConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	instances, err := cfg.ToInstances()
	if err != nil {
		return nil, err
	}

	// Ignore cache error
	SaveNewCache(cfgPath, instances)

	return instances, nil
}
