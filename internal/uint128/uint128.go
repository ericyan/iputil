package uint128

import (
	"encoding/binary"
	"math/big"
)

// Uint128 is a big-endian unsigned 128-bit integer.
type Uint128 struct {
	Hi, Lo uint64
}

// SetBytes interprets buf as the bytes of a big-endian uint128, sets x
// to that value, and returns x.
func (x Uint128) SetBytes(b [16]byte) Uint128 {
	x.Hi = binary.BigEndian.Uint64(b[:8])
	x.Lo = binary.BigEndian.Uint64(b[8:])
	return x
}

// Bytes returns x as a big-endian byte slice.
func (x Uint128) Bytes() [16]byte {
	var buf [16]byte
	binary.BigEndian.PutUint64(buf[:8], x.Hi)
	binary.BigEndian.PutUint64(buf[8:], x.Lo)
	return buf
}

// setDecimalString sets x to the value of s, interpreted in base 10,
// and returns x as well as a boolean indicating success.
func (x Uint128) SetDecimalString(s string) (Uint128, bool) {
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

	return x.SetBytes(b), true
}

// DecimalString converts x to a string of decimal digits.
func (x Uint128) DecimalString() string {
	b := x.Bytes()
	return new(big.Int).SetBytes(b[:]).String()
}
