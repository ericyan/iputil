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

// Pow2 returns 2**n, the base-2 exponential of n.
func Pow2(n uint) (Uint128, error) {
	if n > 127 {
		return Zero, ErrOverflow
	}

	if n < 64 {
		return Uint128{0, 1 << n}, nil
	}

	return Uint128{1 << (n - 64), 0}, nil
}

// Add returns the sum x+y as a new Uint128.
func (x Uint128) Add(y Uint128) Uint128 {
	lo := x.Lo + y.Lo
	hi := x.Hi + y.Hi
	if x.Lo > lo {
		hi++
	}

	return Uint128{hi, lo}
}

// Sub returns the difference x-y as a new Uint128.
func (x Uint128) Sub(y Uint128) Uint128 {
	lo := x.Lo - y.Lo
	hi := x.Hi - y.Hi
	if x.Lo < lo {
		hi--
	}

	return Uint128{hi, lo}
}

// And returns a new Uint128 that is the bitwise AND of two Uint128 values.
func (x Uint128) And(y Uint128) Uint128 {
	return Uint128{x.Hi & y.Hi, x.Lo & y.Lo}
}

// Or returns a new Uint128 that is the bitwise OR of two Uint128 values.
func (x Uint128) Or(y Uint128) Uint128 {
	return Uint128{x.Hi | y.Hi, x.Lo | y.Lo}
}

// Xor returns a new Uint128 that is the bitwise XOR of two Uint128 values.
func (x Uint128) Xor(y Uint128) Uint128 {
	return Uint128{x.Hi ^ y.Hi, x.Lo ^ y.Lo}
}

// Cmp compares x and y and returns either -1, 0, or +1 depending on
// whether x is less than, equal to, or greater than y.
func (x Uint128) Cmp(y Uint128) int {
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

// IsGreaterThan returns true if x is greater than y.
func (x Uint128) IsGreaterThan(y Uint128) bool {
	return x.Cmp(y) == 1
}

// IsLessThan returns true if x is less than y.
func (x Uint128) IsLessThan(y Uint128) bool {
	return x.Cmp(y) == -1
}

// IsEqualTo returns true if x is equal to y.
func (x Uint128) IsEqualTo(y Uint128) bool {
	if x.Hi == y.Hi && x.Lo == y.Lo {
		return true
	}

	return false
}

// IsEven returns true if x is an even number.
func (x Uint128) IsEven() bool {
	return x.And(Uint128{0, 1}) == Zero
}

// IsOdd returns true if x is an odd number.
func (x Uint128) IsOdd() bool {
	return x.And(Uint128{0, 1}) != Zero
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
