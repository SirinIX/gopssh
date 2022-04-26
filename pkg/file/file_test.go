package file

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func TestFilePath(t *testing.T) {
	a := filepath.Join("/", "a", "b", filepath.Base("/c1/c2/c3"))

	fmt.Println(a)
}

func TestGetHome(t *testing.T) {
	u, err := user.Current()
	if err == nil {
		fmt.Println(u.HomeDir)
		return
	}

	envUser := os.Getenv("HOME")
	if envUser == "" {
		fmt.Println("./")
	}
	fmt.Println(envUser)
}

/*
⽣物进化的稳定策略（Evolutionarily StableStrategy，ESS)

ESS是⼀种程序预先编制好的⾏为对策。
这种策略是⼀种抽象情况的对策，⽽不是具体的，因为具体的情况千变万化。
例如“向对⼿进攻；如果它逃就追；如果它还击就逃”就是⼀个策略。


*/
