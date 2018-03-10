package iputil

import (
	"encoding/binary"
	"math/big"
	"net"
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

// ParseDecimal4 parses the string s in base 10 and converts it to IPv4
// address. If s is not valid, ParseDecimal4 returns nil.
func ParseDecimal4(s string) *IPv4 {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return nil
	}

	return ParseUint32(uint32(i.Uint64()))
}

// ParseUint32 parses the integer i as an IPv4 address.
func ParseUint32(i uint32) *IPv4 {
	ip4 := new(IPv4)
	binary.BigEndian.PutUint32(ip4[:], i)
	return ip4
}

// Uint32 converts the IPv4 address to integer.
func (ip4 *IPv4) Uint32() uint32 {
	return binary.BigEndian.Uint32(ip4[:])
}

// String returns the dotted decimal form of the IPv4 address.
func (ip4 *IPv4) String() string {
	return net.IP(ip4[:]).String()
}
