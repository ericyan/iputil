package uint128

import (
	"encoding/binary"
	"errors"
	"math/big"
)

var Zero = Uint128{0x0, 0x0}

// Uint128 is a big-endian unsigned 128-bit integer.
type Uint128 struct {
	Hi, Lo uint64
}

// NewFromBytes creates a new Uint128 from b, a big-endian byte array.
func NewFromBytes(b [16]byte) Uint128 {
	return Uint128{
		Hi: binary.BigEndian.Uint64(b[:8]),
		Lo: binary.BigEndian.Uint64(b[8:]),
	}
}

// NewFromString creates a new Uint128 from s, interpreted in base 10.
func NewFromString(s string) (Uint128, error) {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return Zero, errors.New("invalid string")
	}

	var b [16]byte
	if buf := i.Bytes(); len(buf) > 16 {
		return Zero, errors.New("overflow")
	} else {
		copy(b[16-len(buf):], buf)
	}

	return NewFromBytes(b), nil
}

// Bytes returns x as a big-endian byte slice.
func (x Uint128) Bytes() [16]byte {
	var buf [16]byte
	binary.BigEndian.PutUint64(buf[:8], x.Hi)
	binary.BigEndian.PutUint64(buf[8:], x.Lo)
	return buf
}

// DecimalString converts x to a string of decimal digits.
func (x Uint128) DecimalString() string {
	b := x.Bytes()
	return new(big.Int).SetBytes(b[:]).String()
}
