package config

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestIsIpValidate(t *testing.T) {
	fmt.Println(IsIpValidate("172.18.229"))
}

func TestParseIp(t *testing.T) {
	// 0-255: ([01]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5])
	// (\[\d+:\d+\])|(\d+)
	// pattern := "((\\[[0-255]:[1-255]\\]|[0-255])\\.){3}\\[[0-255]:[1-255]\\]|[0-255]"
	// pattern := "((\\[\\d+:\\d+\\])|(\\d+)\\.){3}(\\[\\d+:\\d+\\])|(\\d+)"
	// pattern := "[01]?[0-9]?[0-9]|2[0-4][0-9]|25[0-5]"
	// pattern := "(\\[\\d+:\\d+\\])|(\\d+)"

	rangePart := "\\[\\d+:\\d+\\]"
	numPart := "\\d+"

	rangeC, err := regexp.Compile(rangePart)
	if err != nil {
		t.Error(err)
	}
	numC, err := regexp.Compile(numPart)
	if err != nil {
		t.Error(err)
	}

	data := []string{
		"[172:173].[16:17].[8:9].[11:13]",
		"172.[16:17].[8:9].[11:13]",
		"172.16.[8:9].[11:13]",
		"172.16.8.[11:13]",
		"999.999.999.k",
	}

	for _, ip := range data {
		ipSplit := strings.Split(ip, ".")
		if len(ipSplit) != 4 {
			t.Errorf("%s is not valid ip", ip)
		}
		ipParts := [][]string{}

		for id, ipPart := range ipSplit {
			if numC.MatchString(ipPart) {
				ipParts[id] = append(ipParts[id], ipPart)
			} else if rangeC.MatchString(ipPart) {
				ipPart = ipPart[1 : len(ipPart)-1]
				ipPartSplit := strings.Split(ipPart, ":")
				if ipPartSplit[1] < ipPartSplit[0] {
					t.Errorf("%s is not valid ip", ip)
				}
				p1, _ := strconv.Atoi(ipPartSplit[0])
				p2, _ := strconv.Atoi(ipPartSplit[1])
				if p1 > p2 {
					t.Errorf("%s is not valid ip", ip)
				}
				for i := p1; i <= p2; i++ {
					ipParts[id] = append(ipParts[id], strconv.Itoa(i))
				}
			} else {
				t.Errorf("%s is not valid ip", ip)
			}
		}
		// ips := []string{}
		// for _, ipPart := range ipParts {
		// 	ip := ""
		// 	for _, ipPartV := range ipPart {
		// 		ip += ipPartV + "."
		// 	}
		// }

	}

}
