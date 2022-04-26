package file

import (
	"fmt"
	"gopssh/log"
	"os"
	"os/user"
)

func OpenFile(path string) (*os.File, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			log.Error("file %s not found, error: %v", path, err)
		}

		log.Error("failed to get file %s status, error: %v", path, err)
		return nil, err
	}

	return os.Open(path)
}

func SaveStringAsFile(path string, content string) error {
	f, err := os.Create(path)
	if err != nil {
		log.Error("failed to create file %s, error: %v", path, err)
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		log.Error("failed to write string to file %s, error: %v", path, err)
		return err
	}

	return nil
}

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

func MustGetUserHome() string {
	userHome := GetUserHome()
	if userHome == "" {
		return "./"
	}

	return userHome
}

func GetUserHome() string {
	u, err := user.Current()
	if err == nil {
		return u.HomeDir
	}

	return os.Getenv("HOME")
}

func EnsureDirExist(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(path, 0755); err != nil {
				log.Error("failed to create dir %s, error: %v", path, err)
				return err
			}
		} else {
			log.Error("failed to get dir %s status, error: %v", path, err)
			return err
		}
	}

	return nil
}

func MustGetFileModTime(path string) int64 {
	stat, _ := os.Stat(path)

	return stat.ModTime().Unix()
}

func GetFileModTime(path string) (int64, error) {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Error("file %s not found, error: %v", path, err)
			return 0, err
		}

		log.Error("failed to get file %s status, error: %v", path, err)
		return 0, err
	}

	return stat.ModTime().Unix(), nil
}
