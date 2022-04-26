package md5

import (
	"crypto/md5"
	"fmt"
)

func Encode(data string) string {
	res := md5.Sum([]byte(data))

	return fmt.Sprintf("%x", res)
}
