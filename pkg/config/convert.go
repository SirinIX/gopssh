package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopssh/log"
	"gopssh/pkg/file"

	"gopkg.in/yaml.v2"
)

const (
	jsonSuffix = ".json"
	yamlSuffix = ".yaml"
)

func UnmarshalConfig(path string) (*Config, error) {
	f, err := file.OpenFile(path)
	if err != nil {
		return nil, err
	}

	if strings.HasSuffix(path, yamlSuffix) {
		return UnmarshalYAML(f)
	} else if strings.HasSuffix(path, jsonSuffix) {
		return UnmarshalJSON(f)
	} else {
		err := fmt.Errorf("only support .json or .yaml config")
		log.Error("unsupported file %v, error: %s", path, err)
	}

	return nil, nil
}

func UnmarshalYAML(f *os.File) (*Config, error) {
	conf := &Config{}

	if err := yaml.NewDecoder(f).Decode(conf); err != nil {
		log.Error("failed to decode yaml file %v, error: %s", f.Name(), err)
		return nil, err
	}

	return conf, nil
}

func UnmarshalJSON(f *os.File) (*Config, error) {
	conf := &Config{}

	if err := json.NewDecoder(f).Decode(conf); err != nil {
		log.Error("failed to decode json file %v, error: %s", f.Name(), err)
		return nil, err
	}

	return conf, nil
}

func (c *Config) ToYAML() (string, error) {
	b, err := yaml.Marshal(c)
	if err != nil {
		log.Error("failed to marshal config to yaml, error: %s", err)
		return "", err
	}

	return string(b), nil
}

func (c *Config) ToJSON() (string, error) {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Error("failed to marshal config to json, error: %s", err)
		return "", err
	}

	return string(b)+"\n", nil
}
