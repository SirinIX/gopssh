package config

import (
	"fmt"
	"gopssh/log"
	"gopssh/pkg/file"
)

const (
	defaultYAMLConfigTemplatePath = "./template.yaml"
	defaultJSONConfigTemplatePath = "./template.json"
)

const (
	defaultYAMLConfigTemplate = `global:
  port: 22
  username: root
  password: cm9vdAo=
  labels:
    all: all
  groups:
  - ips:
    - 192.168.8.8
    - 192.168.8.9
    port: 23
    username: mysql
    password: bXlzcWwK
    labels:
      mysql: master
      # all: all
  - ips:
    - 192.168.8.10
    - 192.168.8.11
    # port: 22
    # username: root
    # password: cm9vdAo=
    labels:
      mysql: slave
      # all: all
`
	defaultJSONConfigTemplate = `{
  "global": {
    "port": 22,
    "username": "root",
    "password": "cm9vdAo=",
    "labels": {
      "all": "all"
	}
  },
  "groups": [
    {
      "ips": [
        "192.168.8.8",
        "192.168.8.9"
      ],
      "port": 23,
      "username": "mysql",
      "password": "bXlzcWwK",
      "labels": {
        "mysql": "master"
      }
    },
    {
      "ips": [
        "192.168.8.10",
        "192.168.8.11"
      ],
      "labels": {
        "mysql": "slave"
      }
    }
  ]
}
`
)

func DumpConfigTemplate(t string) error {
	var cfgStr string
	var filePath string

	// Convert
	if t == "yaml" {
		cfgStr = defaultYAMLConfigTemplate
		filePath = defaultYAMLConfigTemplatePath
	} else if t == "json" {
		cfgStr = defaultJSONConfigTemplate
		filePath = defaultJSONConfigTemplatePath
	} else {
		err := fmt.Errorf("only support yaml or json")
		log.Error("unknown config template type %v, error: %v", t, err)
		return err
	}
	fmt.Printf("The config template is:\n\n%v\n", cfgStr)

	// Save
	if err := file.SaveStringAsFile(filePath, cfgStr); err != nil {
		return err
	}
	fmt.Printf("Successfully saved the config template to %v\n", filePath)

	return nil
}
