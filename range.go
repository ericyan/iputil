package iputil

import (
	"errors"
	"net"

	"github.com/ericyan/iputil/internal/uint128"
)

// A Range represents an arbitrary IP address range.
type Range struct {
	af    uint
	first uint128.Uint128
	last  uint128.Uint128
}

// NewRange returns a new Range.
func NewRange(first, last net.IP) (*Range, error) {
	if AddressFamily(first) != AddressFamily(last) {
		return nil, errors.New("invalid range")
	}
	r := &Range{af: AddressFamily(first)}

	r.first, _ = uint128.NewFromBytes(first)
	r.last, _ = uint128.NewFromBytes(last)
	if !r.first.IsLessThan(r.last) {
		return nil, errors.New("invalid range")
	}

	return r, nil
}

// First returns the first IP address within the range.
func (r *Range) First() net.IP {
	byteLen := 4
	if r.af == IPv6 {
		byteLen = 16
	}

	return r.first.Bytes()[16-byteLen:]
}

// Last returns the last IP address within the range.
func (r *Range) Last() net.IP {
	byteLen := 4
	if r.af == IPv6 {
		byteLen = 16
	}

	return r.last.Bytes()[16-byteLen:]
}

// Contains reports whether the range includes ip.
func (r *Range) Contains(ip net.IP) bool {
	x, _ := uint128.NewFromBytes(ip)

	if x.IsLessThan(r.first) || x.IsGreaterThan(r.last) {
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
	return r.First().String() + " - " + r.Last().String()
}
