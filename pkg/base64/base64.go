package base64

import (
	"encoding/base64"
	"gopssh/log"
)

func Encode(data string) string {
	encStr := base64.StdEncoding.EncodeToString([]byte(data))
	return encStr
}

func Decode(data string) (string, error) {
	decStr, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Error("failed to decode string %v, error: %v", data, err)
		return "", err
	}

	return string(decStr), nil
}
