package iputil

import (
	"math/big"
	"net"
)

// A Range represents an arbitrary IP address range.
type Range struct {
	First net.IP
	Last  net.IP
}

// Contains reports whether the range includes ip.
func (r *Range) Contains(ip net.IP) bool {
	x := new(big.Int).SetBytes(ip)

	if first := new(big.Int).SetBytes(r.First); first.Cmp(x) > 0 {
		return false
	}

	if last := new(big.Int).SetBytes(r.Last); last.Cmp(x) < 0 {
		return false
	}

	return true
}

// Network returns the network name, "ip+net".
func (r *Range) Network() string {
	return "ip+net"
}

// String returns the string form of range.
func (r *Range) String() string {
	return r.First.String() + " - " + r.Last.String()
}
