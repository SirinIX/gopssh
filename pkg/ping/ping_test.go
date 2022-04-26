package ping

import (
	"fmt"
	"net"
	"testing"
	"time"
)

type host struct {
	ip   string
	port int
}

func TestPing(t *testing.T) {
	data := []host{
		{
			"172.16.8.83",
			22,
		},
		{
			"172.16.8.83",
			221,
		},
	}

	for _, ipPort := range data {
		ping(ipPort)
	}
}

func ping(ipPort host) {
	fmt.Println(ipPort)

	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ipPort.ip, ipPort.port), 3*time.Second)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("success")
	}
}
