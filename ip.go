package iputil

import (
	"math/big"
	"net"
)

// AddressFamily returns the address family of given IP address: 4 for
// IPv4, 6 for IPv6, and 0 for everthing else.
func AddressFamily(ip net.IP) uint {
	switch len(ip) {
	case net.IPv4len:
		return 4
	case net.IPv6len:
		return 6
	default:
		return 0
	}
}

// ParseDecimal parses the string s in base 10 and converts it to an IP
// address of specified address family (4 for IPv4 and 6 for IPv6). If
// either s or af is invalid, ParseDecimal returns nil.
func ParseDecimal(s string, af uint) net.IP {
	var byteLen uint
	switch af {
	case 4:
		byteLen = net.IPv4len
	case 6:
		byteLen = net.IPv6len
	default:
		return nil
	}

	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return nil
	}

	ip := make(net.IP, byteLen)
	copy(ip, i.Bytes())

	return ip
}
