package ssh

import (
	"fmt"

	"gopssh/log"
)

type SSH struct {
	Address  *Address      `json:"address" yaml:"address"`
	Username string        `json:"username" yaml:"username"`
	Password string        `json:"password" yaml:"password"`
	Logger   log.CtxLogger `json:"-" yaml:"-"`
}

type Address struct {
	Ip   string `json:"ip" yaml:"ip"`
	Port int    `json:"port" yaml:"port"`
}

type SSHResult struct {
	Addr    string `json:"addr" yaml:"addr"`
	Command string `json:"command" yaml:"command"`
	Stdout  string `json:"stdout" yaml:"stdout"`
	Stderr  string `json:"stderr" yaml:"stderr"`
	// Code    int    `json:"code" yaml:"code"`
}

func (a *Address) String() string {
	return fmt.Sprintf("%s:%d", a.Ip, a.Port)
}
