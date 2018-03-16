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

// IsIPv4 returns true if ip is an 32-bit IPv4 address or an IPv4-mapped
// IPv6 addresses as specified in section 2.5.5.2 of RFC 4291.
func IsIPv4(ip net.IP) bool {
	return AddressFamily(ip.To4()) == 4
}

// IsIPv6 returns true if ip is an 128-bit IPv6 address.
func IsIPv6(ip net.IP) bool {
	return AddressFamily(ip) == 6
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

// DecimalString returns the decimal notation of ip as a string.
func DecimalString(ip net.IP) string {
	return new(big.Int).SetBytes(ip).String()
}
