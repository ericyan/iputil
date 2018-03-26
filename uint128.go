package iputil

import (
	"encoding/binary"
	"math/big"
)

// uint128 is a big-endian unsigned 128-bit integer.
type uint128 [2]uint64

// setBytes interprets buf as the bytes of a big-endian uint128, sets x
// to that value, and returns x.
func (x uint128) setBytes(b [16]byte) uint128 {
	x[0] = binary.BigEndian.Uint64(b[:8])
	x[1] = binary.BigEndian.Uint64(b[8:])
	return x
}

// bytes returns x as a big-endian byte slice.
func (x uint128) bytes() [16]byte {
	var buf [16]byte
	binary.BigEndian.PutUint64(buf[:8], x[0])
	binary.BigEndian.PutUint64(buf[8:], x[1])
	return buf
}

// setDecimalString sets x to the value of s, interpreted in base 10,
// and returns x as well as a boolean indicating success.
func (x uint128) setDecimalString(s string) (uint128, bool) {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return x, false
	}

	var b [16]byte
	if buf := i.Bytes(); len(buf) > 16 {
		return x, false
	} else {
		copy(b[16-len(buf):], buf)
	}

	return x.setBytes(b), true
}

// decimalString converts x to a string of decimal digits.
func (x uint128) decimalString() string {
	b := x.bytes()
	return new(big.Int).SetBytes(b[:]).String()
}
