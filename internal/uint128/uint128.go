package uint128

import (
	"encoding/binary"
	"errors"
	"math/big"
)

var (
	Zero = Uint128{0x0, 0x0}

	ErrOverflow      = errors.New("overflow")
	ErrEmptySlice    = errors.New("empty byte slice")
	ErrInvalidString = errors.New("invalid string")
)

// Uint128 is a big-endian unsigned 128-bit integer.
type Uint128 struct {
	Hi, Lo uint64
}

// NewFromBytes creates a new Uint128 from buf, a big-endian byte slice.
func NewFromBytes(buf []byte) (Uint128, error) {
	if len(buf) == 0 {
		return Zero, ErrEmptySlice
	}

	if len(buf) > 16 {
		return Zero, ErrOverflow
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
		return Zero, ErrInvalidString
	}

	// The zero value for an big.Int represents the value 0.
	if len(i.Bytes()) == 0 {
		return Zero, nil
	}

	return NewFromBytes(i.Bytes())
}

// cmp compares x and y and returns either -1, 0, or +1 depending on
// whether x is less than, equal to, or greater than y.
func (x Uint128) cmp(y Uint128) int {
	if x.Hi > y.Hi {
		return 1
	}
	if x.Hi < y.Hi {
		return -1
	}

	// Now that x.Hi == y.Hi
	if x.Lo > y.Lo {
		return 1
	}
	if x.Lo < y.Lo {
		return -1
	}

	return 0
}

// GreaterThan returns true if x is greater than y.
func (x Uint128) GreaterThan(y Uint128) bool {
	return x.cmp(y) == 1
}

// LessThan returns true if x is less than y.
func (x Uint128) LessThan(y Uint128) bool {
	return x.cmp(y) == -1
}

// EqualTo returns true if x is equal to y.
func (x Uint128) EqualTo(y Uint128) bool {
	if x.Hi == y.Hi && x.Lo == y.Lo {
		return true
	}

	return false
}

// Bytes returns x as a big-endian byte slice.
func (x Uint128) Bytes() []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:8], x.Hi)
	binary.BigEndian.PutUint64(buf[8:], x.Lo)
	return buf
}

// String converts x to a string of decimal digits.
func (x Uint128) String() string {
	return new(big.Int).SetBytes(x.Bytes()).String()
}
