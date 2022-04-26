package ip

import (
	"regexp"
)

const (
	ipPattern     = "^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)$"
	ipPortPattern = "^(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d):\\d+$"
)

var (
	ipRegexp, _     = regexp.Compile(ipPattern)
	ipPortRegexp, _ = regexp.Compile(ipPortPattern)
)

func IsIpValidate(ip string) bool {
	return ipRegexp.MatchString(ip)
}

func IsIpPortValidate(ipPort string) bool {
	return ipPortRegexp.MatchString(ipPort)
}
