package ping

import (
	"net"
	"time"

	"gopssh/log"
	"gopssh/pkg/config"
)

const (
	protocol = "tcp"
	timeout  = 5
)

func Ping(address *config.Address) bool {
	addrStr := address.String()

	_, err := net.DialTimeout(protocol, addrStr, timeout * time.Second)
	if err != nil {
		log.Warning("failed to connect %s, error: %v", addrStr, err)
		return false
	}
	log.Info("succeed to connect %s", addrStr)

	return true
}
