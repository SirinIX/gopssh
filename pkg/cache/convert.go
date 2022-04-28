package cache

import (
	"gopssh/log"
	"gopssh/pkg/file"

	"gopkg.in/yaml.v2"
)

func UnmarshalInstances(path string) (Instances, error) {
	f, err := file.OpenFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal
	instances := &Instances{}
	if err := yaml.NewDecoder(f).Decode(instances); err != nil {
		log.Error("failed to decode yaml file %v, error: %s", f.Name(), err)
		return nil, err
	}

	return *instances, nil
}

func UnmarshalCacheIndex(path string) (*CacheIndex, error) {
	f, err := file.OpenFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal
	index := &CacheIndex{}
	if err := yaml.NewDecoder(f).Decode(index); err != nil {
		log.Error("failed to decode yaml file %v, error: %s", f.Name(), err)
		return nil, err
	}

	return index, nil
}

func (c CacheIndex) ToYAML() (string, error) {
	b, err := yaml.Marshal(c)
	if err != nil {
		log.Error("failed to marshal Cache to yaml, error: %s", err)
		return "", err
	}

	return string(b), nil
}

func (c CacheIndex) ToYAMLBytes() ([]byte, error) {
	b, err := yaml.Marshal(c)
	if err != nil {
		log.Error("failed to marshal Cache to yaml, error: %s", err)
		return nil, err
	}

	return b, nil
}

func (i Instances) ToYAML() (string, error) {
	b, err := yaml.Marshal(i)
	if err != nil {
		log.Error("failed to marshal Instances to yaml, error: %s", err)
		return "", err
	}

	return string(b), nil
}

func (i Instances) ToYAMLBytes() ([]byte, error) {
	b, err := yaml.Marshal(i)
	if err != nil {
		log.Error("failed to marshal Instances to yaml, error: %s", err)
		return nil, err
	}

	return b, nil
}