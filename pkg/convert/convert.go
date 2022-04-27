package convert

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"gopssh/log"

	"github.com/ghodss/yaml"
)

type Convert func (path string) ([]byte, error)

func YAMLFileToJSONBytes(yamlPath string) ([]byte, error) {
	yamlBytes, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Error("failed to read yaml file, error: %s", err)
		return nil, err
	}

	jsonBytes, err := YAMLBytesToJSONBytes(yamlBytes, true)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func JSONFileToYAMLBytes(jsonPath string) ([]byte, error) {
	jsonBytes, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Error("failed to read yaml file, error: %s", err)
		return nil, err
	}

	yamlBytes, err := JSONBytesToYAMLBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return yamlBytes, nil
}

func YAMLBytesToJSONBytes(data []byte, indent bool) ([]byte, error) {
	jsonBytes, err := yaml.YAMLToJSON(data)
	if err != nil {
		log.Error("failed to convert yaml to json, error: %s", err)
		return nil, err
	}

	// Indent json or not
	if !indent {
		return jsonBytes, nil
	}

	var buf bytes.Buffer
	if err := json.Indent(&buf, jsonBytes, "", "  "); err != nil {
		log.Error("failed to indent json, error: %s", err)
		return nil, err
	}

	return buf.Bytes(), nil
}

func JSONBytesToYAMLBytes(data []byte) ([]byte, error) {
	yamlBytes, err := yaml.JSONToYAML(data)
	if err != nil {
		log.Error("failed to convert json to yaml, error: %s", err)
		return nil, err
	}

	return yamlBytes, nil
}
