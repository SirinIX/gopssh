package port

import (
	"net"
	"time"

	"gopssh/log"
	"gopssh/pkg/cache"
)

const (
	protocol = "tcp"
	timeout  = 5
)

func CheckPort(address *cache.Address) bool {
	addrStr := address.String()

	_, err := net.DialTimeout(protocol, addrStr, timeout * time.Second)
	if err != nil {
		log.Debug("failed to connect %s, error: %v", addrStr, err)
		return false
	}

	return true
}
