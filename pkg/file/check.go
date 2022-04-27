package file

import (
	"fmt"
	"gopssh/log"
	"os"
)

func IsPathExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func IsPathExistE(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			err := fmt.Errorf("config file not found")
			log.Error("file %v is not exist, error: %v", path, err)
			return err
		}
	}

	return nil
}
