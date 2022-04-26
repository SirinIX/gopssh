package config

import (
	"gopssh/pkg/file"
	"path/filepath"
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
	Labels   map[string]string `json:"labels" yaml:"labels"`
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
		Labels: map[string]string{
			"all": "all",
		},
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
