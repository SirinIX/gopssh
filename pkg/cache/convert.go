package cache

import (
	"gopssh/log"
	"gopssh/pkg/file"

	"gopkg.in/yaml.v2"
)

func UnmarshalCache(path string) (*Cache, error) {
	f, err := file.OpenFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal
	conf := &Cache{}
	if err := yaml.NewDecoder(f).Decode(conf); err != nil {
		log.Error("failed to decode yaml file %v, error: %s", f.Name(), err)
		return nil, err
	}

	return conf, nil
}

func (c *Cache) ToYAML() (string, error) {
	b, err := yaml.Marshal(c)
	if err != nil {
		log.Error("failed to marshal Cache to yaml, error: %s", err)
		return "", err
	}

	return string(b), nil
}
