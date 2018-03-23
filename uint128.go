package iputil

import (
	"encoding/binary"
)

// uint128 is a big-endian unsigned 128-bit integer.
type uint128 [2]uint64

// setBytes interprets buf as the bytes of a big-endian uint128, sets x
// to that value, and returns x.
func (x uint128) setBytes(b []byte) uint128 {
	x[0] = binary.BigEndian.Uint64(b[:8])
	x[1] = binary.BigEndian.Uint64(b[8:])
	return x
}

// bytes returns x as a big-endian byte slice.
func (x uint128) bytes() []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:8], x[0])
	binary.BigEndian.PutUint64(buf[8:], x[1])
	return buf
}
