package config

import (
	"fmt"
	"gopssh/log"
	"gopssh/pkg/base64"
	"gopssh/pkg/cache"
	pkgIp "gopssh/pkg/ip"
)

func (c *Config) ToInstances() (cache.Instances, error) {
	var err error
	var instances cache.Instances

	for _, group := range c.Groups {
		group.CombineGlobalSetting(c.Global)
		// Decode password
		group.Password, err = base64.Decode(group.Password)
		if err != nil {
			return nil, err
		}
		// Build instance
		for _, ip := range group.Ips {
			// Check ip
			if !pkgIp.IsIpValidate(ip) {
				err := fmt.Errorf("invalid ip")
				log.Error("ip %v is not a valid ip, error: %v", ip, err)
				return nil, err
			}

			inst := &cache.Instance{
				Address: &cache.Address{
					Ip:   ip,
					Port: group.Port,
				},
				Username: group.Username,
				Password: group.Password,
				Labels:   group.Labels,
			}
			instances = append(instances, inst)
		}
	}

	return instances, nil
}

func (g *Group) CombineGlobalSetting(global *Global) {
	if g.Port == 0 {
		g.Port = global.Port
	}
	if g.Username == "" {
		g.Username = global.Username
	}
	if g.Password == "" {
		g.Password = global.Password
	}

	if g.Labels == nil {
		g.Labels = global.Labels
	} else {
		for k, v := range global.Labels {
			g.Labels[k] = v
		}
	}
}
