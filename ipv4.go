package iputil

import (
	"net"
	"strconv"
)

type IPv4 [4]byte

// ParseIPv4 parses the string s as an IPv4 address. If s is not a valid
// textual representation of an IPv4 address, ParseIPv4 returns nil.
func ParseIPv4(s string) *IPv4 {
	ip := net.ParseIP(s).To4()
	if ip == nil {
		return nil
	}

	ip4 := new(IPv4)
	copy(ip4[:], ip)
	return ip4
}

// String returns the dotted decimal form of the IPv4 address.
func (ip4 *IPv4) String() string {
	return strconv.Itoa(int(ip4[0])) + "." +
		strconv.Itoa(int(ip4[1])) + "." +
		strconv.Itoa(int(ip4[2])) + "." +
		strconv.Itoa(int(ip4[3]))
}
