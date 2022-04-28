package rand

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphanNums = "bcdfghjklmnpqrstvwxz2456789"
)

func String(length int) string {
	if length < 1 {
		return ""
	}
	charArr := strings.Split(alphanNums, "")
	charlen := len(charArr)

	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var res string = ""
	for i := 1; i <= length; i++ {
		res = res + charArr[ran.Intn(charlen)]
	}

	return res
}
