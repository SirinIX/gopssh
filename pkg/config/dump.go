package config

import (
	"fmt"
	"gopssh/log"
	"gopssh/pkg/file"
)

const (
	defaultYAMLConfigTemplateName = "template.yaml"
	defaultJSONConfigTemplateName = "template.json"
)

var configTemplate = &Config{
	Global: &Global{
		Port:     22,
		Username: "root",
		Password: "cm9vdAo=",
		Labels: map[string]string{
			"all": "all",
		},
	},
	Groups: Groups{
		&Group{
			Ips: []string{
				"192.168.8.8",
				"192.168.8.9",
			},
			Port:     23,
			Username: "mysql",
			Password: "bXlzcWwK",
			Labels: Labels{
				"mysql": "master",
			},
		},
		&Group{
			Ips: []string{
				"192.168.8.10",
				"192.168.8.11",
			},
			Labels: Labels{
				"mysql": "slave",
			},
		},
	},
}

func DumpConfigTemplate(t string) error {
	var err error
	var cfgStr string
	var fileName string

	// Convert
	if t == "yaml" {
		cfgStr, err = configTemplate.ToYAML()
		if err != nil {
			return err
		}
		fileName = defaultYAMLConfigTemplateName
	} else if t == "json" {
		cfgStr, err = configTemplate.ToJSON()
		if err != nil {
			return err
		}
		fileName = defaultJSONConfigTemplateName
	} else {
		err := fmt.Errorf("only support yaml or json")
		log.Error("unknown config template type %v, error: %v", t, err)
		return err
	}
	fmt.Printf("The config template is:\n%v\n", cfgStr)

	// Save
	if err := file.SaveStringAsFile("./"+fileName, cfgStr); err != nil {
		return err
	}
	fmt.Printf("Successfully saved the config template to %v\n", fileName)

	return nil
}
