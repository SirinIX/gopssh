package ssh

import (
	"fmt"
	"strings"
)

type SSH struct {
	Address  *Address `json:"address" yaml:"address"`
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
	// Logger   *log.CtxLogger `json:"-" yaml:"-"`
}

type Address struct {
	Ip   string `json:"ip" yaml:"ip"`
	Port int    `json:"port" yaml:"port"`
}

type SSHResult struct {
	Address *Address `json:"address" yaml:"address"`
	Command string   `json:"command" yaml:"command"`
	Stdout  string   `json:"stdout" yaml:"stdout"`
	Stderr  string   `json:"stderr" yaml:"stderr"`
	// Code    int    `json:"code" yaml:"code"`
}

func (a *Address) String() string {
	return fmt.Sprintf("%s:%d", a.Ip, a.Port)
}

func (s *SSHResult) String() string {
	stdout := strings.TrimSpace(s.Stdout)
	if stdout != "" {
		stdout = "\n" + stdout
	}
	stderr := strings.TrimSpace(s.Stderr)
	if stderr != "" {
		stderr = "\n" + stderr
	}
	return fmt.Sprintf("[ Host ]: %s\n[ Command ]: %s\n[ Stdout ]: %s\n[ Stderr ]: %s", s.Address.String(), s.Command, stdout, stderr)
}