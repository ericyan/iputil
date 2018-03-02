package iputil

import (
	"math/big"
	"net"
)

type IPv6 [16]byte

// ParseIPv6 parses the string s as an IPv6 address. If s is not a valid
// textual representation of an IPv6 address, ParseIPv6 returns nil.
func ParseIPv6(s string) *IPv6 {
	ip := net.ParseIP(s).To16()
	if ip == nil {
		return nil
	}

	ip6 := new(IPv6)
	copy(ip6[:], ip)
	return ip6
}

// ParseBigInt parses the integer i as an IPv6 address.
func ParseBigInt(i *big.Int) *IPv6 {
	ip6 := new(IPv6)
	copy(ip6[:], i.Bytes())
	return ip6
}

// BigInt converts the IPv6 address to integer.
func (ip6 *IPv6) BigInt() *big.Int {
	i := big.NewInt(0)
	return i.SetBytes(ip6[:])
}

// String returns the dotted decimal form of the IPv6 address.
func (ip6 *IPv6) String() string {
	return net.IP(ip6[:]).String()
}
