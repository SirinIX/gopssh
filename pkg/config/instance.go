package config

import "fmt"

type Instance struct {
	Address  *Address `json:"address" yaml:"address"`
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
	Labels   Labels   `json:"labels" yaml:"labels"`
}

type Address struct {
	Ip   string `json:"ip" yaml:"ip"`
	Port int    `json:"port" yaml:"port"`
}

func (a *Address) String() string {
	return fmt.Sprintf("%s:%d", a.Ip, a.Port)
}
