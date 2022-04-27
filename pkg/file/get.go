package file

import (
	"gopssh/log"
	"os"
	"os/user"
)

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
