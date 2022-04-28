package file

import (
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
			log.Error("file %v is not exist, error: %v", path, err)
			return err
		}
		log.Error("failed to get file %v info, error: %v", path, err)
		return err
	}

	return nil
}
