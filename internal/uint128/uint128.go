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

// NewFromBytes creates a new Uint128 from buf, a big-endian byte slice.
func NewFromBytes(buf []byte) (Uint128, error) {
	if len(buf) == 0 {
		return Zero, errors.New("empty byte slice")
	}

	if len(buf) > 16 {
		return Zero, errors.New("overflow")
	}

	var b [16]byte
	copy(b[16-len(buf):], buf)

	return Uint128{
		Hi: binary.BigEndian.Uint64(b[:8]),
		Lo: binary.BigEndian.Uint64(b[8:]),
	}, nil
}

// NewFromString creates a new Uint128 from s, interpreted in base 10.
func NewFromString(s string) (Uint128, error) {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return Zero, errors.New("invalid string")
	}

	// The zero value for an big.Int represents the value 0.
	if len(i.Bytes()) == 0 {
		return Zero, nil
	}

	return NewFromBytes(i.Bytes())
}

// Bytes returns x as a big-endian byte slice.
func (x Uint128) Bytes() []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:8], x.Hi)
	binary.BigEndian.PutUint64(buf[8:], x.Lo)
	return buf
}

// DecimalString converts x to a string of decimal digits.
func (x Uint128) DecimalString() string {
	return new(big.Int).SetBytes(x.Bytes()).String()
}
