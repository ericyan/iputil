package iputil

import (
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

// ParseIPv4 parses s as an IPv4 address.
func ParseIPv4(s string) net.IP {
	return net.ParseIP(s).To4()
}

// ParseIPv6 parses s as an IPv6 address. If s is an IPv4 address, it
// will be converted to a IPv4-mapped IPv6 address.
func ParseIPv6(s string) net.IP {
	return net.ParseIP(s).To16()
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

	x, ok := new(uint128).setDecimalString(s)
	if !ok {
		return nil
	}

	ip := make(net.IP, byteLen)
	b := x.bytes()
	copy(ip, b[16-byteLen:])

	return ip
}

// DecimalString returns the decimal notation of ip as a string. If ip
// has length 0, it returns "<nil>".
func DecimalString(ip net.IP) string {
	if len(ip) == 0 {
		return "<nil>"
	}

	var b [16]byte
	copy(b[16-len(ip):], ip)

	return new(uint128).setBytes(b).decimalString()
}
