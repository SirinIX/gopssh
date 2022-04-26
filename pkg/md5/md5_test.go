package md5

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	data := `version = "0.0.1"`
	res := md5.Sum([]byte(data))

	resStr := fmt.Sprintf("%x", res)

	fmt.Println(resStr)
}