package iputil

import (
	"net"

	"github.com/ericyan/iputil/internal/uint128"
)

// A Range represents an arbitrary IP address range.
type Range struct {
	First net.IP
	Last  net.IP
}

// Contains reports whether the range includes ip.
func (r *Range) Contains(ip net.IP) bool {
	x, _ := uint128.NewFromBytes(ip)

	if first, _ := uint128.NewFromBytes(r.First); x.IsLessThan(first) {
		return false
	}

	if last, _ := uint128.NewFromBytes(r.Last); x.IsGreaterThan(last) {
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
