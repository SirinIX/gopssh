package config

type Config struct {
	Global Global `json:"global"`
	Groups Groups `json:"groups"`
}

type Global struct {
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Labels   Labels `json:"labels"`
}

type Groups []Group

type Group struct {
	Ips      []string `json:"ips"`
	Port     int      `json:"port"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Labels   Labels   `json:"labels"`
}

type Labels map[string]string

func NewGlobal() *Global {
	return &Global{
		Port:     22,
		Username: "root",
		Labels: map[string]string{
			"all": "all",
		},
	}
}
